# XDCHELP

This repository contains a simple read-only function from a proxy smart contract, which already works on the Polygon Mainnet and Amoy. To keep the environment as close as possible to our real environment, we use Docker to maintain the same build configurations.

To run the script, you must have Docker installed on your machine, and the Daemon must be running.

To build the image, in the root directory, type:

```bash
docker build -t image_name .
```

To run the created image, type:

```bash
docker run --rm -it image_name
```

Replace `image_name` with the name you want to give your Docker image. The `--rm` option ensures the container is automatically removed after it stops, and `-it` makes the terminal interactive, allowing you to interact with the container during its execution.
