import logo from './logo.svg';
import './App.css';
import { useQuery } from '@apollo/client';
import gql from 'graphql-tag';
import { RoomsQuery } from './generated/graphql';

function Rooms() {

  const query = gql`
    query Rooms {
      rooms {
        id
        players {
          name
        }
        voting
      }
    }
  `

  const { loading, error, data } = useQuery<RoomsQuery>(query)


  if (loading || !data) return <p>Loading...</p>;

  if (error) return <p>Error :(</p>;

  const view = data.rooms.map(
    room =>
    (
      <p key={room.id}>
        Room: {room.id}
      </p>
    )
  )

  return <>{view}</>

}

function App() {
  return (
    <>
      <Rooms></Rooms>
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    </>
  );
}

export default App;
