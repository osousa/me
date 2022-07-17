BINARY_NAME=me
OUTPUT_DIR=${CURDIR}/output

livematrix:
	echo "Building..."

build:
	cd web && ./tailwind.sh && cd ..
	rm -rf ${OUTPUT_DIR} && mkdir -p ${OUTPUT_DIR} && cp ${CURDIR}/.env ${OUTPUT_DIR}/.env
	mkdir -p ${OUTPUT_DIR}/web && cp -R ${CURDIR}/web/html ${OUTPUT_DIR}/web/html
	cp -R ${CURDIR}/web/static ${OUTPUT_DIR}/web/static
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ${OUTPUT_DIR}/${BINARY_NAME} .
	# upx --best --lzma ${OUTPUT_DIR}/${BINARY_NAME} 

all: livematrix build 
