URL crontab 
===

Lightweight implementation of HTTP call focused crontab. Can be used for GET and POST calls with custom HTTP headers and 
data payload.

## Configuration format

Schedule is specified as regular crontab. Content can be provided either as mounted file crontab in project root or as 
string via environment variable `CRONTAB`. In case file is mounted and environment variable `CRONTAB` provided - result
crontab is combination of both.

Cron time options consist of 7 positions: second, minute, hour, day of month, month, day of week and year.

```shell
#   +-------------------- Second 0-59
#   | +------------------ Minute 0-59
#   | | *---------------- Hour 0-23
#   | | | +-------------- Day of month 1-31
#   | | | | +------------ Month 1-12
#   | | | | | +---------- Day of week 0-6 (o Sunday)
#   | | | | | | +-------- Year
#   | | | | | | |
#   * * * * * * *         <HTTP request specification>
```

It is possible to use shortened version with only 5 time options where second and year is omitted. 

Crontab has modified - URL centric structure. After time options comes URL target definition in the following format:

```shell
{url} {method} {headers} {payload} 
```

- `{url}` HTTP URL with schema
- `{method}` HTTP method. "GET" and "POST" are currently supported. If not provided - "GET" is default
- `{headers}` HTTP headers in form `<key:value>`. Might be added as much as needed.
- `{payload}` Payload as string which is sent if HTTP method is "POST"

Header values can contain sensitive data which can be provided via environment variables in order to avoid exposition 
in code repository. Reference to environment variable is defined as `${VAR_NAME}`.

Example of HTTP POST request every 15 minutes with content-type and authorization headers and data payload:

` * */15 * * * * * https://google.com POST <content-type:application/json> <authorization: bearer ${TOKEN}> {"id":1}`

## Telemetry

Service provide telemetry in Prometheus format which can be obtained via http request `http://<servicename>:80/metrics`.  

## Running as Docker container

When running service as container configuration can be supplied either via mount or environment variable, or both. 
Container is available in dockerhub.

Example when config in `/home/user/crontab`:

```shell
docker run --rm -v /home/user/crontab:/usr/src/app/crontab saleniex/cronurl
```

## Configuration

Application support following environment variables:

- `ADDR` - Address and port where http service is listening. Default any IP address on port `80`.
- `CRONTAB` - Crontab definition. Might be multiline separated with `\n` (multiline). 
- `VERBOSE` - Enable verbose behaviour. All scheduler activities are logged. Default `false`.
