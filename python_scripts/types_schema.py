import tomllib
import json

# Attempt to parse
with open("schema.toml", "rb") as f:
    try:
        parsed = tomllib.load(f)
        data = parsed
        print(json.dumps(data, indent=2, sort_keys=True))
    except tomllib.TOMLDecodeError as e:
        print("TOML error:", e)
