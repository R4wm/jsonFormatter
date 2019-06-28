# jsonFormatter
format json file

```bash
ᚱ@jsonFormatter $
ᚱ@jsonFormatter $ echo '{"haha": "yourface"}' > /tmp/stuff.json
ᚱ@jsonFormatter $ echo '{"haha": "yourface"}' > /tmp/stuff1.json
ᚱ@jsonFormatter $
ᚱ@jsonFormatter $ go run main.go /tmp/stuff.json /tmp/stuff1.json
ᚱ@jsonFormatter $
ᚱ@jsonFormatter $
ᚱ@jsonFormatter $ cat /tmp/stuff*.json
{
  "haha": "yourface"
}{
  "haha": "yourface"
}ᚱ@jsonFormatter $

```

I know, there's lots to fix here, but its just for me so.. maybe later?