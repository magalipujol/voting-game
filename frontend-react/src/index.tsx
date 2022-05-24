import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  useQuery,
  gql,
} from "@apollo/client";

const client = new ApolloClient({
  uri: "http://localhost:8080/graphql",
  cache: new InMemoryCache(),
});

let f = async () => {
  let res = await client.mutate({
    mutation: gql`
      mutation CreatePlayer {
        createPlayer(name: "agus") {
          id
        }
      }
    `,
  });
  let playerId = res.data.createPlayer.id;

  res = await client.mutate({
    mutation: gql`
      mutation CreateRoom {
        createRoom(name: "test") {
          id
        }
      }
    `,
  });
  let roomId = res.data.createRoom.id;

  client.mutate({
    mutation: gql`
    mutation JoinRoom {
    joinRoom(playerId: "${playerId}", roomId: "${roomId}") {
        id
    }
    }
`,
  });

  client
    .query({
      query: gql`
        query GetRooms {
          rooms {
            id
            players {
              name
            }
            voting
          }
        }
      `,
    })
    .then((result) => console.log(result.data));
};

f();

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
