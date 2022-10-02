#!/bin/bash
printf "%s$(ls -l | sed -n 'n;p')"