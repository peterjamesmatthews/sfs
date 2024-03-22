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
    "\nquery FileByID($id: ID) {\n  node(id: $id) {\n    ... on File {\n      name\n      owner { id }\n      parent { id }\n      content\n    }\n  }\n}\n": types.FileByIdDocument,
    "\nquery FolderByID($id: ID) {\n  node(id: $id) {\n    ... on Folder {\n      name\n      owner { id }\n      parent { id }\n      children { id }\n    }\n  }\n}\n": types.FolderByIdDocument,
    "\nquery Node {\n  node {\n    ... on Folder {\n      children {\n        id\n      }\n    }\n  }\n}\n": types.NodeDocument,
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
export function gql(source: "\nquery FileByID($id: ID) {\n  node(id: $id) {\n    ... on File {\n      name\n      owner { id }\n      parent { id }\n      content\n    }\n  }\n}\n"): (typeof documents)["\nquery FileByID($id: ID) {\n  node(id: $id) {\n    ... on File {\n      name\n      owner { id }\n      parent { id }\n      content\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery FolderByID($id: ID) {\n  node(id: $id) {\n    ... on Folder {\n      name\n      owner { id }\n      parent { id }\n      children { id }\n    }\n  }\n}\n"): (typeof documents)["\nquery FolderByID($id: ID) {\n  node(id: $id) {\n    ... on Folder {\n      name\n      owner { id }\n      parent { id }\n      children { id }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery Node {\n  node {\n    ... on Folder {\n      children {\n        id\n      }\n    }\n  }\n}\n"): (typeof documents)["\nquery Node {\n  node {\n    ... on Folder {\n      children {\n        id\n      }\n    }\n  }\n}\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;