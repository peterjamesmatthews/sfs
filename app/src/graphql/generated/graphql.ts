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


export type MutationCreateUserArgs = {
  name: Scalars['String']['input'];
  password: Scalars['String']['input'];
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
  Me: User;
  getTokens: Tokens;
};


export type QueryGetTokensArgs = {
  name: Scalars['String']['input'];
  password: Scalars['String']['input'];
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

export type CreateUserMutationVariables = Exact<{
  name: Scalars['String']['input'];
  password: Scalars['String']['input'];
}>;


export type CreateUserMutation = { __typename?: 'Mutation', createUser?: { __typename?: 'User', id: string, name: string } | null };

export type GetTokensQueryVariables = Exact<{
  name: Scalars['String']['input'];
  password: Scalars['String']['input'];
}>;


export type GetTokensQuery = { __typename?: 'Query', getTokens: { __typename?: 'Tokens', access: string, refresh: string } };

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', Me: { __typename?: 'User', name: string } };

export type RefreshTokensMutationVariables = Exact<{
  refresh: Scalars['String']['input'];
}>;


export type RefreshTokensMutation = { __typename?: 'Mutation', refreshTokens?: { __typename?: 'Tokens', access: string, refresh: string } | null };


export const CreateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"password"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}},{"kind":"Argument","name":{"kind":"Name","value":"password"},"value":{"kind":"Variable","name":{"kind":"Name","value":"password"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CreateUserMutation, CreateUserMutationVariables>;
export const GetTokensDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetTokens"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"password"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getTokens"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}},{"kind":"Argument","name":{"kind":"Name","value":"password"},"value":{"kind":"Variable","name":{"kind":"Name","value":"password"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"access"}},{"kind":"Field","name":{"kind":"Name","value":"refresh"}}]}}]}}]} as unknown as DocumentNode<GetTokensQuery, GetTokensQueryVariables>;
export const MeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Me"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Me"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<MeQuery, MeQueryVariables>;
export const RefreshTokensDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"RefreshTokens"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"refresh"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"refreshTokens"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"refresh"},"value":{"kind":"Variable","name":{"kind":"Name","value":"refresh"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"access"}},{"kind":"Field","name":{"kind":"Name","value":"refresh"}}]}}]}}]} as unknown as DocumentNode<RefreshTokensMutation, RefreshTokensMutationVariables>;