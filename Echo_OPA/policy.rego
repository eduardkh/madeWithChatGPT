package myapp.authz

default allow = false

# Allow access if the user role is "admin"
allow {
	input.user == "admin"
}
