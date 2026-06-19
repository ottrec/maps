# maps

Map styles for ottrec.ca. WIP.

Edit these in Maputnik, then import the changes manually. The OSM data inspector is useful (`Layers > Map Data`).

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
  - [ ] Look at raw OSM extracts from https://download.geofabrik.de/.
  - [ ] Convert them to mvt or mlt with https://github.com/systemed/tilemaker or https://github.com/versatiles-org/versatiles-rs or https://github.com/felt/tippecanoe.
  - [ ] Render them with https://github.com/maptiler/tileserver-gl or something else?
- [ ] Figure out how to render them.
