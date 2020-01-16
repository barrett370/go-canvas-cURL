url=$1
auth=$2
output=$3

curl --location --request GET $url --header 'Authorization: ${auth}' --output $output