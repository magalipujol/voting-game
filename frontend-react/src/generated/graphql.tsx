import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Mutation = {
  __typename?: 'Mutation';
  createPlayer: Player;
  createRoom: Room;
  deletePlayer?: Maybe<Player>;
  deleteRoom?: Maybe<Room>;
  joinRoom: Room;
  leaveRoom?: Maybe<Room>;
  startGame: Room;
  vote: Room;
};


export type MutationCreatePlayerArgs = {
  name: Scalars['String'];
};


export type MutationCreateRoomArgs = {
  name: Scalars['String'];
};


export type MutationDeletePlayerArgs = {
  id: Scalars['ID'];
};


export type MutationDeleteRoomArgs = {
  id: Scalars['ID'];
};


export type MutationJoinRoomArgs = {
  playerId: Scalars['ID'];
  roomId: Scalars['ID'];
};


export type MutationLeaveRoomArgs = {
  id: Scalars['ID'];
};


export type MutationStartGameArgs = {
  roomId: Scalars['ID'];
};


export type MutationVoteArgs = {
  option: Scalars['String'];
  playerId: Scalars['ID'];
};

export type Option = {
  __typename?: 'Option';
  value: Scalars['String'];
};

export type Player = {
  __typename?: 'Player';
  id: Scalars['ID'];
  name: Scalars['String'];
  room: Room;
};

export type Query = {
  __typename?: 'Query';
  room?: Maybe<Room>;
  rooms: Array<Room>;
};


export type QueryRoomArgs = {
  id: Scalars['ID'];
};

export type Room = {
  __typename?: 'Room';
  id: Scalars['ID'];
  options: Array<Option>;
  players: Array<Player>;
  votes: Array<Vote>;
  voting: Scalars['Boolean'];
};

export type Vote = {
  __typename?: 'Vote';
  option: Option;
  player: Player;
  room: Room;
  timeStamp: Scalars['String'];
};

export type RoomsQueryVariables = Exact<{ [key: string]: never; }>;


export type RoomsQuery = { __typename?: 'Query', rooms: Array<{ __typename?: 'Room', id: string, voting: boolean, players: Array<{ __typename?: 'Player', name: string }> }> };

export type CreatePlayerMutationVariables = Exact<{ [key: string]: never; }>;


export type CreatePlayerMutation = { __typename?: 'Mutation', createPlayer: { __typename?: 'Player', id: string } };

export type CreateRoomMutationVariables = Exact<{ [key: string]: never; }>;


export type CreateRoomMutation = { __typename?: 'Mutation', createRoom: { __typename?: 'Room', id: string } };

export type JoinRoomMutationVariables = Exact<{ [key: string]: never; }>;


export type JoinRoomMutation = { __typename?: 'Mutation', joinRoom: { __typename?: 'Room', id: string } };

export type GetRoomsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetRoomsQuery = { __typename?: 'Query', rooms: Array<{ __typename?: 'Room', id: string, voting: boolean, players: Array<{ __typename?: 'Player', name: string }> }> };


export const RoomsDocument = { "kind": "Document", "definitions": [{ "kind": "OperationDefinition", "operation": "query", "name": { "kind": "Name", "value": "Rooms" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "rooms" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "id" } }, { "kind": "Field", "name": { "kind": "Name", "value": "players" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "name" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "voting" } }] } }] } }] } as unknown as DocumentNode<RoomsQuery, RoomsQueryVariables>;
export const CreatePlayerDocument = { "kind": "Document", "definitions": [{ "kind": "OperationDefinition", "operation": "mutation", "name": { "kind": "Name", "value": "CreatePlayer" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "createPlayer" }, "arguments": [{ "kind": "Argument", "name": { "kind": "Name", "value": "name" }, "value": { "kind": "StringValue", "value": "agus", "block": false } }], "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "id" } }] } }] } }] } as unknown as DocumentNode<CreatePlayerMutation, CreatePlayerMutationVariables>;
export const CreateRoomDocument = { "kind": "Document", "definitions": [{ "kind": "OperationDefinition", "operation": "mutation", "name": { "kind": "Name", "value": "CreateRoom" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "createRoom" }, "arguments": [{ "kind": "Argument", "name": { "kind": "Name", "value": "name" }, "value": { "kind": "StringValue", "value": "test", "block": false } }], "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "id" } }] } }] } }] } as unknown as DocumentNode<CreateRoomMutation, CreateRoomMutationVariables>;
export const JoinRoomDocument = { "kind": "Document", "definitions": [{ "kind": "OperationDefinition", "operation": "mutation", "name": { "kind": "Name", "value": "JoinRoom" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "joinRoom" }, "arguments": [{ "kind": "Argument", "name": { "kind": "Name", "value": "playerId" }, "value": { "kind": "StringValue", "value": "", "block": false } }, { "kind": "Argument", "name": { "kind": "Name", "value": "roomId" }, "value": { "kind": "StringValue", "value": "", "block": false } }], "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "id" } }] } }] } }] } as unknown as DocumentNode<JoinRoomMutation, JoinRoomMutationVariables>;
export const GetRoomsDocument = { "kind": "Document", "definitions": [{ "kind": "OperationDefinition", "operation": "query", "name": { "kind": "Name", "value": "GetRooms" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "rooms" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "id" } }, { "kind": "Field", "name": { "kind": "Name", "value": "players" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "name" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "voting" } }] } }] } }] } as unknown as DocumentNode<GetRoomsQuery, GetRoomsQueryVariables>;
