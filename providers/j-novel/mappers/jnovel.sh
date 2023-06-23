#!/bin/sh

exec curl 'https://labs.j-novel.club/app/v1/events?sort=launch&start_date=2023-05-31T21%3A00%3A00.000Z&end_date=2023-06-30T21%3A00%3A00.000Z' --compressed \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/114.0' \
-H 'Accept: */*' \
-H 'Accept-Language: en-US,en;q=0.5' \
-H 'Accept-Encoding: gzip, deflate' \
-H 'Referer: https://j-novel.club/' \
-H 'Content-Type: application/json' \
-H 'Origin: https://j-novel.club' \
-o response.proto
