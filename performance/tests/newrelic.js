import http from 'k6/http';

export class NewRel {
    constructor(apiKey, log = false) {
        this.log = log;
        this.params = {
            headers: {
                'Content-Type': 'application/json',
                'API-Key': apiKey,
            },
        };
        this.urls = {
            graphql: 'https://api.newrelic.com/graphql',
        };
    }

    PrintAlertingStatus() {
        let payload = JSON.stringify({
            query: `
      {
        actor {
          entitySearch(query: "name LIKE 'Node Workshop' AND domain IN ('APM')") {
            results {
              entities {
                ... on ApmApplicationEntityOutline {
                    alertSeverity
                }
              }
            }
          }
        }
      }
      `,
        });

        let res = http.post(this.urls.graphql, payload, this.params);

        if (this.log) {
            console.log('New Relic Check HTTP Status is: ' + res.status);
        }

        if (res.status === 200) {
            let body = JSON.parse(res.body);
            var result = body.data.actor.entitySearch.results.entities[0].alertSeverity;
            console.log('New Relic Status: ' + result);
        }
    }

    AppID() {

        // From NerdGraph, copy the GraphQL payload from tools > copy as cURL > take the entire {"query"} section
        let payload = JSON.stringify({
            query: `
      {
        actor {
          entitySearch(query: "name LIKE 'Node Workshop' AND domain IN ('APM')") {
            results {
              entities {
                ... on ApmApplicationEntityOutline {
                  applicationId
                }
              }
            }
          }
        }
      }
    `
        });

        let res = http.post(this.urls.graphql, payload, this.params);
        // Check we are not experiencing HTTP 400. If you are, the payload is likely wrong.

        if (this.log) {
            console.log('New Relic Check HTTP Status is: ' + res.status);
        }


        if (res.status === 200) {
            let body = JSON.parse(res.body);
            /* result will depend on the query. This query is built on alertSeverity result.
             You need to modify the selector if you are performing a different query     */
            var result = JSON.stringify(
                body.data.actor.entitySearch.results.entities[0].applicationId
            );
        } else {
            throw new Error('Could not fetch AppID from New Relic')
        }

        return result;
    }

    //Send a deployment marker with start/end information on load test.
    Notify(testName, state, description, user) {
        var url =
            'https://api.newrelic.com/v2/applications/' + this.AppID() + '/deployments.json';
        console.log(url);

        // From NerdGraph, copy the GraphQL payload from tools > copy as cURL > take the entire {"query"} section
        let payload = JSON.stringify({
            deployment: {
                revision: testName,
                changelog: 'k6 load test ' + state,
                description: description,
                user: user,
            },
        });

        let res = http.post(url, payload, this.params);
        // Check we are not experiencing HTTP 400. If you are, the payload is likely wrong.
        if (![200, 201].includes(res.status)) {
            throw new Error(`Could not notify New Relic about test state (res: ${res.status})`)
        }

        return JSON.stringify(res.status);
    }
}