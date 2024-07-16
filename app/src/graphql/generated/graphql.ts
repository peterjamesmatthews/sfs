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
  moveNode?: Maybe<Node>;
  refreshTokens?: Maybe<Tokens>;
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


export type MutationMoveNodeArgs = {
  id: Scalars['ID']['input'];
  parentID?: InputMaybe<Scalars['ID']['input']>;
};


export type MutationRefreshTokensArgs = {
  refresh: Scalars['String']['input'];
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
  getNodeFromPath?: Maybe<Node>;
  getTokensFromAuth0: Tokens;
  me: User;
};


export type QueryGetNodeFromPathArgs = {
  path: Scalars['String']['input'];
};


export type QueryGetTokensFromAuth0Args = {
  token: Scalars['String']['input'];
};

export type Tokens = {
  __typename?: 'Tokens';
  access: Scalars['String']['output'];
  refresh: Scalars['String']['output'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type GetNodeFromPathQueryVariables = Exact<{
  path: Scalars['String']['input'];
}>;


export type GetNodeFromPathQuery = { __typename?: 'Query', getNodeFromPath?: { __typename?: 'File', content: string, id: string, name: string, parent?: { __typename?: 'Folder', id: string, name: string } | null } | { __typename?: 'Folder', id: string, name: string, children: Array<{ __typename?: 'File', id: string, name: string } | { __typename?: 'Folder', id: string, name: string }>, parent?: { __typename?: 'Folder', id: string, name: string } | null } | null };

export type GetTokensFromAuth0QueryVariables = Exact<{
  token: Scalars['String']['input'];
}>;


export type GetTokensFromAuth0Query = { __typename?: 'Query', getTokensFromAuth0: { __typename?: 'Tokens', access: string, refresh: string } };

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me: { __typename?: 'User', name: string } };

export type RefreshTokensMutationVariables = Exact<{
  refresh: Scalars['String']['input'];
}>;


export type RefreshTokensMutation = { __typename?: 'Mutation', refreshTokens?: { __typename?: 'Tokens', access: string, refresh: string } | null };


export const GetNodeFromPathDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetNodeFromPath"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"path"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getNodeFromPath"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"path"},"value":{"kind":"Variable","name":{"kind":"Name","value":"path"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"parent"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Folder"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"File"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"content"}}]}}]}}]}}]} as unknown as DocumentNode<GetNodeFromPathQuery, GetNodeFromPathQueryVariables>;
export const GetTokensFromAuth0Document = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetTokensFromAuth0"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"token"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getTokensFromAuth0"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"token"},"value":{"kind":"Variable","name":{"kind":"Name","value":"token"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"access"}},{"kind":"Field","name":{"kind":"Name","value":"refresh"}}]}}]}}]} as unknown as DocumentNode<GetTokensFromAuth0Query, GetTokensFromAuth0QueryVariables>;
export const MeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Me"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"me"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<MeQuery, MeQueryVariables>;
export const RefreshTokensDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"RefreshTokens"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"refresh"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"refreshTokens"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"refresh"},"value":{"kind":"Variable","name":{"kind":"Name","value":"refresh"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"access"}},{"kind":"Field","name":{"kind":"Name","value":"refresh"}}]}}]}}]} as unknown as DocumentNode<RefreshTokensMutation, RefreshTokensMutationVariables>;