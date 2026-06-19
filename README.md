# maps

Map styles for ottrec.ca. WIP.

Edit these in Maputnik, then import the changes manually. The OSM data inspector is useful (`Layers > Map Data`).

### Rendering

Get the latest North America OSM extract from [Geofabrik](https://download.geofabrik.de/).

Create OpenMapTiles-compatible vector tiles with [tilemaker](https://github.com/systemed/tilemaker).

```bash
docker run --interactive --rm -v $(pwd):/data ghcr.io/systemed/tilemaker:master /data/north-america-latest.osm.pbf --output /data/ottrec.pmtiles --bbox -76.6,44.8,-75.0,45.7 --config /data/tilemaker.json
```

The map bounds are `44.8,-76.6` to `45.7,-75.0`, and also note that `45.5,-76.2` to `45.2,-75.1` contains all facilities with some extra room.

Even though the zoom is set to 14 for the vector tiles, we can render the raster ones at a higher zoom level.

Due to https://github.com/systemed/tilemaker/issues/802, the labels in the styles need to use `{name:latin}` instead of Carto's `{name}` or `{name_en}`.

To test these in Maputnik, use `"carto": { "type": "vector", "tiles": ["http://localhost:8080/ottrec/{z}/{x}/{y}.mvt"] }` with `go run github.com/protomaps/go-pmtiles@v1.30.3 serve tilemaker --cors='*' --public-url=http://localhost:8080`.

TODO: figure out how to render raster tiles

### TODO

#### Features

- [ ] Remove unneeded features to simplify customization.
  - [X] Airports
  - [X] Boundary lines
  - [ ] More?
- [ ] Remove clutter
  - [X] Wading pool labels (these are a bit silly, it's marking every little puddle of water, and some of the swimming pools)
  - [X] Park labels
  - [ ] More?
- [ ] Add transit stuff
  - [X] LRT
  - [ ] LRT stations
  - [ ] Bus stops

#### Misc

- [ ] Tweak colors to match site.
- [ ] See if I can color LRT lines differently.
- [ ] ~~See if OpenMapTiles datasets from https://www.maptiler.com/on-prem-datasets/planet/ are sufficient.~~ They aren't, and they aren't free either.
  - [X] Look at raw OSM extracts from https://download.geofabrik.de/.
  - [X] Convert them to mvt or mlt with https://github.com/systemed/tilemaker ~~or https://github.com/versatiles-org/versatiles-rs or https://github.com/felt/tippecanoe~~.
  - [ ] Render them with https://github.com/maptiler/tileserver-gl or https://github.com/ConservationMetrics/mapgl-tile-renderer or something else?
- [ ] Figure out how to render them.
