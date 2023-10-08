# Feature Flags Service

It's a Feature Flag lightweight micro-service.
Feature Flags are cool, micro-services are cool; that makes this project is cool as well :/

I created this project after struggling with feature flags in one of my applications.
Then I used LaunchDarkly but, I had some issues there.

FFS is a replacement option for my in-app embedded stuff, LaunchDarkly or other similar services.

One can create and update flags and users via API or WebUI or directly via SQL.

Literally tens of thousands of active users can be served from a $5/mo VPS.

## Dependencies

This software was created with Go (1.10), HTML/CSS/JS, Riot (3) PostgreSQL (10), Redis (3)

The primary Go server is built (via `make`) to serve some static HTML and handle some API requests.
You should not expose this system to the wider internet without putting some type of security in front of it.
You can use Nginx or other reverse proxys, adjust firewall rules, etc.

Feature Flags are stored in PostgreSQL and both Flags and Users are cached in Redis.

## Installing

Simply clone this repo, and say `make`.
Then create the database container, using an existing PostgreSQL instance or a new one (this service is very lightweight).
Configure using the `ffs.ini` file.
Run.

## Using the API

### Create a Flag

	POST /api/v2016/flag
	name=Some%20Name&type=bool

	POST /api/v2016/flag
	name=Some%20Name&type=enum&enum=Thing%20One,Thing%20Two


### Get a User and their Flags at once

When fetching a user provide a HASH, such as sha1.
This should be created outside of FFS and could be based on your application groups, or user or UID.
We just need to use a consistent one each time.

	GET /api/v2016/user/{hash}

If a User object doesn't exist one will be magically created and the response code is 201 rather than 200.

### Update a User

	POST /api/v2016/user/{hash}
	name=Some%20Name

