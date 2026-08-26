# Cursor blog feed provider

> :information_source: Minimum Viable Implementation
>
> Limitations are listed based on the last 50 posts fetched from the
> Cursor blog on 2026-08-26.

## Known limitations

- Splash images/videos are very inconsistent in the way they are supplied so
  OpenGraph data is used to maximise compatibility with RSS readers.
- Alt parameter for splash image is hard-coded as it's not always present in
  the page content (especially if splash is a video).
- SVG graphs that the team loves to use, are styled via bundled (not always
  inlined) CSS. Fetching CSS bundles and replaying the logic that web browsers
  use for rendering looks like quite an overkill (for the
  brave souls who would like to fix this, contributions are welcome, I just send such posts to Instapaper)
