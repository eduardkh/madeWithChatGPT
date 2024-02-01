import requests
from bs4 import BeautifulSoup
import pandas as pd
from io import StringIO

url = "https://crt.sh/?q=crt.sh"
response = requests.get(url)


def format_key(key):
    # Replace spaces and special characters with underscores
    return key.replace(" ", "_").replace(".", "_").lower()


if response.status_code == 200:
    soup = BeautifulSoup(response.content, 'html.parser')
    table = soup.find_all('table')[1]

    # Wrap the HTML in StringIO
    html_content = StringIO(str(table))
    df = pd.read_html(html_content)[1]

    # Format column names
    df.columns = [format_key(col) for col in df.columns]

    # Save to JSON as a list of objects
    df.to_json("output.json", orient="records", lines=False)
else:
    print("Error fetching data")
