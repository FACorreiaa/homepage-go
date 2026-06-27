---
title: "MarkusPfundstein/mcp-obsidian: MCP server that interacts with Obsidian via the Obsidian rest API community plugin"
source: "https://github.com/MarkusPfundstein/mcp-obsidian"
author:
published:
created: 2026-04-25
description: "MCP server that interacts with Obsidian via the Obsidian rest API community plugin - MarkusPfundstein/mcp-obsidian"
tags:
  - "clippings"
---
## MCP server for Obsidian

MCP server to interact with Obsidian via the Local REST API community plugin.

[![server for Obsidian MCP server](https://camo.githubusercontent.com/bd8277adb242b2869aeda4b0c71f0946ef29a5f2265330b988fc0d57a2c2496d/68747470733a2f2f676c616d612e61692f6d63702f736572766572732f33776b6f31626875656b2f6261646765)](https://glama.ai/mcp/servers/3wko1bhuek)

## Components

### Tools

The server implements multiple tools to interact with Obsidian:

- list\_files\_in\_vault: Lists all files and directories in the root directory of your Obsidian vault
- list\_files\_in\_dir: Lists all files and directories in a specific Obsidian directory
- get\_file\_contents: Return the content of a single file in your vault.
- search: Search for documents matching a specified text query across all files in the vault
- patch\_content: Insert content into an existing note relative to a heading, block reference, or frontmatter field.
- append\_content: Append content to a new or existing file in the vault.
- delete\_file: Delete a file or directory from your vault.

### Example prompts

Its good to first instruct Claude to use Obsidian. Then it will always call the tool.

The use prompts like this:

- Get the contents of the last architecture call note and summarize them
- Search for all files where Azure CosmosDb is mentioned and quickly explain to me the context in which it is mentioned
- Summarize the last meeting notes and put them into a new note 'summary meeting.md'. Add an introduction so that I can send it via email.

## Configuration

### Obsidian REST API Key

There are two ways to configure the environment with the Obsidian REST API Key.

1. Add to server config (preferred)
```
{
  "mcp-obsidian": {
    "command": "uvx",
    "args": [
      "mcp-obsidian"
    ],
    "env": {
      "OBSIDIAN_API_KEY": "<your_api_key_here>",
      "OBSIDIAN_HOST": "<your_obsidian_host>",
      "OBSIDIAN_PORT": "<your_obsidian_port>"
    }
  }
}
```

Sometimes Claude has issues detecting the location of uv / uvx. You can use `which uvx` to find and paste the full path in above config in such cases.

1. Create a `.env` file in the working directory with the following required variables:

```
OBSIDIAN_API_KEY=your_api_key_here
OBSIDIAN_HOST=your_obsidian_host
OBSIDIAN_PORT=your_obsidian_port
```

Note:

- You can find the API key in the Obsidian plugin config
- Default port is 27124 if not specified
- Default host is 127.0.0.1 if not specified

## Quickstart

### Install

#### Obsidian REST API

You need the Obsidian REST API community plugin running: [https://github.com/coddingtonbear/obsidian-local-rest-api](https://github.com/coddingtonbear/obsidian-local-rest-api)

Install and enable it in the settings and copy the api key.

#### Claude Desktop

On MacOS: `~/Library/Application\ Support/Claude/claude_desktop_config.json`

On Windows: `%APPDATA%/Claude/claude_desktop_config.json`

Development/Unpublished Servers Configuration
```
{
  "mcpServers": {
    "mcp-obsidian": {
      "command": "uv",
      "args": [
        "--directory",
        "<dir_to>/mcp-obsidian",
        "run",
        "mcp-obsidian"
      ],
      "env": {
        "OBSIDIAN_API_KEY": "<your_api_key_here>",
        "OBSIDIAN_HOST": "<your_obsidian_host>",
        "OBSIDIAN_PORT": "<your_obsidian_port>"
      }
    }
  }
}
```
Published Servers Configuration
```
{
  "mcpServers": {
    "mcp-obsidian": {
      "command": "uvx",
      "args": [
        "mcp-obsidian"
      ],
      "env": {
        "OBSIDIAN_API_KEY": "<YOUR_OBSIDIAN_API_KEY>",
        "OBSIDIAN_HOST": "<your_obsidian_host>",
        "OBSIDIAN_PORT": "<your_obsidian_port>"
      }
    }
  }
}
```

## Development

### Building

To prepare the package for distribution:

1. Sync dependencies and update lockfile:
```
uv sync
```

### Debugging

Since MCP servers run over stdio, debugging can be challenging. For the best debugging experience, we strongly recommend using the [MCP Inspector](https://github.com/modelcontextprotocol/inspector).

You can launch the MCP Inspector via [`npm`](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) with this command:

```
npx @modelcontextprotocol/inspector uv --directory /path/to/mcp-obsidian run mcp-obsidian
```

Upon launching, the Inspector will display a URL that you can access in your browser to begin debugging.

You can also watch the server logs with this command:

```
tail -n 20 -f ~/Library/Logs/Claude/mcp-server-mcp-obsidian.log
```