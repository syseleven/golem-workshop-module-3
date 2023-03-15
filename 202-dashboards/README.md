# Grafana

## Dashboards

### Scope

**This installation is required only once per cluster.**

* Since Grafana is available we need to install some more dashboards.

### Inspect

* Inspect the prepared dashboard manifests in the `dashboards` folder.

* Take note of the `labels: grafana_dashboard: "1"` flag which instructs Grafana sidecar
container to pick up these dashboard definitions and apply them at runtime.

### Deploy dashboards

* Now deploy the additional dashboards:

  ```shell
  kubectl -n monitoring apply -f dashboards/
  ```

---

* Optional - only if dashboards are not loaded automatically

  ```shell
  # Force reloading of grafana dashboards if needed
  kubectl rollout restart deployment prom-grafana
  ```

---

### View the result

* Verify the changed in the Grafana web interface:
  * new dashboards are present
    * blackbox-exporter
    * nginx

### Task

* Use the "Explore" view and query builder in Grafana to obtain metrics from our application.

* Example:

  `sum(ping_request_count)`
