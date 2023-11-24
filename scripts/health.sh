until pg_isready -U postgres; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done