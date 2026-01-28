# Python Examples

This directory contains Python examples demonstrating how to use Zaphiro Protocol Buffers with RabbitMQ.

## Examples

- **measurements/** - Produce and consume measurement data using RabbitMQ Streams
- **faults/** - Produce and consume fault events using RabbitMQ AMQP

## Installation

This project uses Poetry for dependency management. Install dependencies with:

```bash
cd examples/python
poetry install
```

## Running Examples

After installing dependencies, you can run the examples using poetry:

```bash
# Run measurements example
poetry run python measurements/main.py

# Run faults example
poetry run python faults/main.py
```

Or activate the virtual environment and run directly:

```bash
poetry shell
python measurements/main.py
python faults/main.py
```

## Dependencies

- **rstream** (>=0.15.0) - RabbitMQ Streams client
- **pika** (>=1.3.0) - RabbitMQ AMQP client
- **protobuf** (^5.29.5) - Protocol Buffers runtime

Dependencies are managed via `pyproject.toml` and locked in `poetry.lock`.
