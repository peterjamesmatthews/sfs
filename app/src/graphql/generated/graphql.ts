/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Access = {
  __typename?: 'Access';
  target: Node;
  type: AccessType;
  user: User;
};

export enum AccessType {
  Read = 'READ',
  Write = 'WRITE'
}

export type File = Node & {
  __typename?: 'File';
  content: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  owner: User;
  parent?: Maybe<Folder>;
};

export type Folder = Node & {
  __typename?: 'Folder';
  children: Array<Node>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  owner: User;
  parent?: Maybe<Folder>;
};

export type Mutation = {
  __typename?: 'Mutation';
  createFile?: Maybe<File>;
  createFolder?: Maybe<Folder>;
  createUser?: Maybe<User>;
  moveNode?: Maybe<Node>;
  renameNode?: Maybe<Node>;
  shareNode?: Maybe<Access>;
  writeFile?: Maybe<File>;
};


export type MutationCreateFileArgs = {
  content?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  parentID?: InputMaybe<Scalars['ID']['input']>;
};


export type MutationCreateFolderArgs = {
  name: Scalars['String']['input'];
  parentID?: InputMaybe<Scalars['ID']['input']>;
};


export type MutationCreateUserArgs = {
  name: Scalars['String']['input'];
};


export type MutationMoveNodeArgs = {
  id: Scalars['ID']['input'];
  parentID?: InputMaybe<Scalars['ID']['input']>;
};


export type MutationRenameNodeArgs = {
  id: Scalars['ID']['input'];
  name: Scalars['String']['input'];
};


export type MutationShareNodeArgs = {
  accessType: AccessType;
  targetID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type MutationWriteFileArgs = {
  content: Scalars['String']['input'];
  id: Scalars['ID']['input'];
};

export type Node = {
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  owner: User;
  parent?: Maybe<Folder>;
};

export type Query = {
  __typename?: 'Query';
  me: User;
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me: { __typename?: 'User', id: string, name: string } };


export const MeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Me"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"me"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<MeQuery, MeQueryVariables>;