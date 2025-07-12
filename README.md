# prometheus-shelly-update-checker

A simple prometheus probe/exporter for checking whether a Shelly
device needs an update. Check Prometheus configuration below.

```yaml
  - job_name: shelly
    scrape_interval: 1h
    metrics_path: /probe
    static_configs:
         - targets:
            - 'http://kitchen-dimmer.iot.hemma/status'
            - 'http://office-ceiling.iot.hemma/rpc/Sys.GetStatus'
            - 'http://office-led.iot.hemma/status'
            - 'http://outside-back.iot.hemma/status'
            - 'http://laundry-ceiling.iot.hemma/status'
    relabel_configs:
         - source_labels: [__address__]
           target_label: __param_target
         - source_labels: [__param_target]
           target_label: instance
         - target_label: __address__
           replacement: docker.hemma:7980
```