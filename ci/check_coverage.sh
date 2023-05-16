#!/bin/sh

COVERAGE_STRING=$(go tool cover -func coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')

COVERAGE_PASS=$(echo "$COVERAGE_STRING >= 90.0" | bc)

EXIT_CODE=$(echo "1 - $COVERAGE_PASS" | bc)

echo "Coverage is $COVERAGE_STRING"

exit $EXIT_CODE