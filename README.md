# Go Tiled Development Server

I find this useful for quickly prototyping new levels in web based games,
instead of having to re-export to `.png` constantly.

## Setup

Place all `.tmx` map files in the `/maps` directory. The
development server will serve these as static assets. For any file
`/maps/my_map.tmx` and request for `/maps/my_map.png` will return
the proper `.png` encoding of that map file. Ensure that the `Accept`
header in these cases is set to `image/png`.