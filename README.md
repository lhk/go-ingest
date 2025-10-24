## Building

`go build -o go-ingest`

## Usage
The executable takes an arbitrary number of file patterns as command line parameters.
Here I'm piping the output to a file, to not flood the terminal

`./go-ingest lklein/go-ingest/*.md lklein/go-ingest/*.go > output.txt`

output.txt will look like this
`cat output.txt | head -20

# Folder structure
└── lklein
    ├── cachesaver
    │   ├── API_proposal.md
    │   ├── README.md
    │   └── prompts.md
    ├── dueling_agents
    │   └── README.md
    ├
    ├── go-ingest
    │   └── README.md
