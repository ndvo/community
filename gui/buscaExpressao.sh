#!/bin/bash

ack --ignore-dir=dist-prod --ignore-dir=dist --ignore-file=match:/.*\.json/ --ignore-file=match:/.*\.go/ --ignore-file=match:/.*\.css/ --ignore-file=match:/.*\.map/  $1 $2
