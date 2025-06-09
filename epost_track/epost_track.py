import argparse
import re
import requests
import sys


def normalize_id(raw_id):
    match = re.search(r'(\d+)', raw_id)
    return match.group(1) if match else None


def track_shipment(epost_id):
    epost_id = normalize_id(epost_id)

    if not epost_id:
        print("Invalid Epost ID")
        sys.exit(1)

    # Use a placeholder HFD value, not important
    dummy_hfd = "000000000"

    url = f"https://run.hfd.co.il/RunCom.Server/Request.aspx?APPNAME=RUN&PRGNAME=ship_locate_blank&arguments=-AEP{epost_id},-AHfd{dummy_hfd},-Ay,-Ajson"

    headers = {
        'User-Agent': 'Mozilla/5.0',
        'Accept': 'application/json',
        'Referer': 'https://run.hfd.co.il/EPOST_TRACK/',
    }

    try:
        response = requests.get(url, headers=headers)
        response.raise_for_status()
        data = response.json().get("root", {})
    except Exception as e:
        print(f"Failed to fetch or parse response: {e}")
        sys.exit(1)

    print(f"\n📦 Shipment Number: {data.get('shipment_num')}")
    print(f"👤 Receiver: {data.get('reciever_name')}")
    print(f"🏢 Sender: {data.get('company_name')}")
    print(f"📍 Pickup Point: {data.get('name_pudo')} ({data.get('address')})")
    print(f"🕐 Opening Hours: {data.get('comments_pudo')}")
    print(f"📞 Driver: {data.get('driver_name')} ({data.get('driver_phone')})")
    print("\n📜 Tracking History:")

    for stage in data.get("lines", []):
        print(
            f" - {stage['taarich_shlav']} {stage['shaa_shlav']} - {stage['teur_shlav_mishloah']}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Track Epost shipment using Epost ID only")
    parser.add_argument(
        "epost_id", help="Epost ID (with or without 'EP' prefix)")

    args = parser.parse_args()
    track_shipment(args.epost_id)
