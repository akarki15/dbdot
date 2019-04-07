#!/usr/bin/env bash
# Credit: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
# see https://stackoverflow.com/questions/38952282/bad-array-subscript-while-splitting-a-string-in-bash
# package_name=${package_split[-1]}
package_name="dbdot"

platforms=(
 "darwin/386"
 "darwin/amd64"
 "freebsd/386"
 "freebsd/amd64"
 "linux/386"
 "linux/amd64"
 "linux/ppc64"
 "linux/ppc64le"
 "linux/mips"
 "linux/mipsle"
 "linux/mips64"
 "linux/mips64le"
 "netbsd/386"
 "netbsd/amd64"
 "openbsd/386"
 "openbsd/amd64"
 "windows/386"
 "windows/amd64"
)

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='release/'$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    echo 'Compiling for '$GOOS' '$GOARCH' to '$output_name

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
