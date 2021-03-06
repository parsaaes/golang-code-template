package config

const Namespace = "golang_code_template"

//nolint:lll
const Default = `
logger:
  level: debug

server:
  address: :7677
  read-timeout: 20s
  write-timeout: 20s
  graceful-timeout: 5s

redis:
  master:
    address: 127.0.0.1:6379
    pool-size: 0
    min-idle-conns: 20
    dial-timeout: 5s
    read-timeout: 3s
    write-timeout: 3s
    pool-timeout: 4s
    idle-timeout: 5m
    max-retries: 5
    min-retry-backoff: 1s
    max-retry-backoff: 3s
  slave:
    address: 127.0.0.1:6379
    pool-size: 0
    min-idle-conns: 20
    dial-timeout: 5s
    read-timeout: 3s
    write-timeout: 3s
    pool-timeout: 4s
    idle-timeout: 5m
    max-retries: 5
    min-retry-backoff: 1s
    max-retry-backoff: 3s

postgres:
  host: 127.0.0.1
  port: 54320
  user: template
  pass: secret
  dbname: template
  connect-timeout: 30s
  connection-lifetime: 30m
  max-open-connections: 100
  max-idle-connections: 10

monitoring:
  prometheus:
    enabled: true
    address: ":9001"
`
