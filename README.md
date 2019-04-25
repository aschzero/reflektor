# reflektor

reflektor is a utility for scheduled directory or Docker volume archival for periodic backups.


## Configuration

Jobs are defined in a yaml configuration file with a name, source path to be archived, and the scheduled defined using cron.

Example:

```yaml
jobs:
  - name: prometheus
    source: /prometheus
    schedule: "@monthly"

  - name: elasticsearch
    source: /elasticsearch
    schedule: "@weekly"
    
  - name: unifi_backups
    source: /unifi
    schedule: "0 0 * * SUN"
```

## Usage

### Command Line

Build reflektor using `make build`:

```bash
cd /path/to/reflektor
make build
./reflektor -config.file=/path/to/config.yml
```

### Using Docker

Mount the directories to be archived along with your config file in a reflektor container.
Ensure archives are saved to the host machine with a mount pointing to `/archives`:

```bash
docker run -v /path/to/config.yml:/config.yml \
    -v prometheus_prometheus:/prometheus \
    -v logs_elasticsearch:/elasticsearch \
    -v unifi_data:/unifi \
    -v /host/path/to/archives:/archives \
    aschzero/reflektor:latest -config.file=/config.yml
```

## Archives

When a job is ready to run, an archive of the source directory is archived in a `.tar.gz` file in the `/archives` directory.

## Logging

 When started, reflektor parses the config file and registers each job as a cron schedule. Registered job details are logged
 to stdout details including the next run time.
 
 ```bash
INFO[0000] job registered       job=prometheus          next_run="2019-05-01 00:00:00 -0700 PDT"
INFO[0000] job registered       job=elasticsearch       next_run="2019-04-28 00:00:00 -0700 PDT"
INFO[0000] job registered       job=unifi_backups       next_run="2019-04-27 00:00:00 -0700 PDT"
```

Logs for finished running jobs include the elapsed time it took for archival:

```bash
INFO[0003] job running        job=prometheus
INFO[0003] job finished       elapsed=225.157555s job=prometheus next_run="2019-06-01 00:00:00 -0700 PDT"
```
