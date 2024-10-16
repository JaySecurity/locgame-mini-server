#!/bin/sh

# If the user has supplied only arguments append them to `locg-service` command
if [ "${1:0:1}" = '-' ]; then
	set -- locg-service  "$@"
fi

exec "$@"