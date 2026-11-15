# Sky

A Go CLI for observed aircraft, nearby traffic, and regional airline activity using [adsb.lol](https://www.adsb.lol/).

## Build

Requires Go 1.24 or newer.

```sh
go build -buildvcs=false -o bin/sky ./cmd/sky
./bin/sky --help
```

## Commands

```sh
# Aircraft registration, or an explicit ICAO hex address
sky plane G-VBEL
sky plane --hex 4073ce --json

# Broadcast callsign; this is not a scheduled flight-number lookup
sky callsign BAW123

# Coordinates are required; radius is in kilometres (default 100, max 463)
sky nearby --lat 6.58 --lon 3.32 --radius-km 100
sky nearby --lat 51.47 --lon -0.45 --airborne --json

# Observed operator activity within an area, using a three-letter ICAO prefix
sky airline BAW --lat 51.47 --lon -0.45 --radius-km 100 --airborne

# Optional provider API key; no key is embedded in the binary
export SKY_API_KEY='your-key'
sky plane G-VBEL --timeout 10s
```

If running the binary directly after building, replace `sky` above with `./bin/sky`, or put it on your PATH.

Each invocation makes one request. There is no automatic polling or retry loop. HTTP 429 errors surface the provider's Retry-After header. Consult the provider's current quota before automating requests; a once-per-minute query uses 1,440 requests/day.

## Reading results

- Table output includes registration, broadcast callsign, aircraft type, barometric altitude (feet), ground speed (knots), ground track (degrees), coordinates, message age, and position age.
- `ground` is the provider's explicit ground altitude value. `--airborne` includes numeric barometric altitude observations and excludes unknown/ground altitude; it is not authoritative flight status.
- A dash means unavailable. Zero is preserved as a valid reading.
- Message and position ages are seconds before the provider snapshot, not the local fetch time. Both snapshot and retrieval timestamps are included.
- JSON wraps normalized provider fields in `source`, `retrieved_at`, `snapshot_at`, and `aircraft`. Altitude may be numeric or `"ground"`; missing readings remain absent. An empty result has `"aircraft": []`.
- No observation does not establish whether an aircraft is landed, inactive, or outside receiver coverage.
- Airline filtering uses broadcast callsign prefixes. It is a regional view of observed operator activity, not an authoritative airline fleet or codeshare lookup.

## Scope

Implemented: aircraft and callsign lookup, nearby search, regional operator filtering, JSON output, timeout and rate-limit handling, interrupt cancellation.

Next: persisted watchlists and quota-aware polling; SQLite observation history; a separate flight-status provider for dated flight numbers, schedules, routes, delays, and ETAs. No hosted backend is required for the current commands.

## Development

```sh
go test ./...
go vet ./...
```

Tests use local HTTP servers and require no provider credentials or live network.

Data source: [adsb.lol API](https://www.adsb.lol/docs/open-data/api), published under ODbL 1.0. Review the provider's license and current access limits before redistributing data.
