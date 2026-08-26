#!/bin/sh

set -eu

host="${FEEDHUB_HOST:-127.0.0.1}"
port="${PORT:-8080}"

base="http://${host}:${port}"
outdir="${1:-.}"

mkdir -p "${outdir}"

curl -fsS "${base}/anthropic/news.atom" -o "${outdir}/anthropic-news.xml"
curl -fsS "${base}/anthropic/engineering.atom" -o "${outdir}/anthropic-engineering.xml"
curl -fsS "${base}/anthropic/research.atom" -o "${outdir}/anthropic-research.xml"
curl -fsS "${base}/anthropic/research/frontier-red-team.atom" -o "${outdir}/anthropic-research-frontier-red-team.xml"
