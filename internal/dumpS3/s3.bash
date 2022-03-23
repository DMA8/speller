#!/bin/bash
URL="https://s3-svc-dl.wildberries.ru/"
BUCKET="s3://spell_check"
FILE_TO_DUMP="./Dump/spellcheck.csv"

# crontab -e add line  "* */6 * * * ./s3.bash"

echo connecting to $URL
# aws  --endpoint-url=$URL s3 ls
# aws  --endpoint-url=$URL s3 mb $BUCKET
aws  --endpoint-url=$URL s3 mv $FILE_TO_DUMP $BUCKET
echo finished
