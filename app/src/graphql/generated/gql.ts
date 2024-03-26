/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\nquery GetFileByID($id: ID!) {\n  getFileById(id: $id) {\n    name\n    owner { id }\n    parent { id }\n    content\n  }\n}\n": types.GetFileByIdDocument,
    "\nquery GetFolderByID($id: ID!) {\n  getFolderById(id: $id) {\n    ... on Folder {\n      name\n      owner { id }\n      parent { id }\n      children { id name }\n    }\n  }\n}\n": types.GetFolderByIdDocument,
    "\nquery GetNodeByURI($uri: String!) {\n  getNodeByURI(uri: $uri) {\n    id\n    name\n    owner { id }\n    parent { id }\n  }\n}\n": types.GetNodeByUriDocument,
    "\nquery GetRoot {\n  getRoot {\n    children {\n      id\n    }\n  }\n}\n": types.GetRootDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery GetFileByID($id: ID!) {\n  getFileById(id: $id) {\n    name\n    owner { id }\n    parent { id }\n    content\n  }\n}\n"): (typeof documents)["\nquery GetFileByID($id: ID!) {\n  getFileById(id: $id) {\n    name\n    owner { id }\n    parent { id }\n    content\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery GetFolderByID($id: ID!) {\n  getFolderById(id: $id) {\n    ... on Folder {\n      name\n      owner { id }\n      parent { id }\n      children { id name }\n    }\n  }\n}\n"): (typeof documents)["\nquery GetFolderByID($id: ID!) {\n  getFolderById(id: $id) {\n    ... on Folder {\n      name\n      owner { id }\n      parent { id }\n      children { id name }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery GetNodeByURI($uri: String!) {\n  getNodeByURI(uri: $uri) {\n    id\n    name\n    owner { id }\n    parent { id }\n  }\n}\n"): (typeof documents)["\nquery GetNodeByURI($uri: String!) {\n  getNodeByURI(uri: $uri) {\n    id\n    name\n    owner { id }\n    parent { id }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery GetRoot {\n  getRoot {\n    children {\n      id\n    }\n  }\n}\n"): (typeof documents)["\nquery GetRoot {\n  getRoot {\n    children {\n      id\n    }\n  }\n}\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;