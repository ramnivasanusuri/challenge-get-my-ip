# Challenge - Get My IP

#### Files provided:

```
├── challenge-get-my-iP
│   ├── README.md
│   ├── input_fqdn.csv
│   ├── sample_output.csv
│   ├── index.html
│   └── server.go
```

- How to use files:
  - `input_fqdn.csv` use as input to your code
  - `sample_output.csv` sample of the desired output formatting

#### Coding Tasks

- For each FQDN, in 'input_fqdn.csv`, extract the following values:
  | field | description |
  |---------|--------------|
  | IP | The IPv4 address rendered from a DNS query |

- create a output.csv file similar to the sample_output.csv

- Use only Go Programming Language, to code your solution.
- In your solution code: try to handle all the possible input types and errors.
- Hint: You may need to import some Network Programming Modules

#### Code-Design Concepts to Consider

- Pay attention to ways you would manage and improve your code for:
  - Scalability (when list has 50,000+ websites)
  - Reliability
  - Speed
  - Testability

#### Extra Credit

- The Server.Go must render your output file at port 8090.
- Additional credit if you can render the output csv in a tabular format in the browser
- Give details on how you would manage dependency requirements.

#### Submission of Work

- Submit/commit your code as _new_ files in your personal repo folder
  - Name your repo folder: `challenge-get-my-ip`
- Please provide complete written instructions on how to use your solution files.
  - e.g. Like how to run, build, deploy, etc.

## Solution

???
