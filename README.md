# ftsb
ftsb is a Discord bot that gives automatic updates regarding the status of trails in Fredericksburg.

---

## Installation w/ Docker (Recommended)
```shell
git clone https://github.com/devhsoj/ftsb
cd ftsb/
docker build . -t ftsb
```
Running:
```shell
docker run -d -e DISCORD_BOT_TOKEN='xxxx' ftsb
```

---

### Installation w/o Docker
Requirements: [Go](https://go.dev/doc/install)
```shell
git clone https://github.com/devhsoj/ftsb
cd ftsb/
go build .
```
Running Inline:
```shell
DISCORD_BOT_TOKEN='xxxx' ./ftsb
```
Running with a `.env` or with a system/user environment variable:
```text
DISCORD_BOT_TOKEN='xxxx'
```
then
```shell
./ftsb
```

---

### ftsb commands
**Note:** ftsb will automatically send the trail status summary every 4 hours by default.

---

To get the trail status summary:
```text
!trailstatus
```