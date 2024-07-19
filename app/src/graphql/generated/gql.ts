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
    "\nquery GetNodeFromPath($path: String!) {\n  getNodeFromPath(path: $path) {\n    id\n    name\n    parent {\n      id\n      name\n    }\n    ...on Folder {\n      children {\n        id\n        name\n      }\n    }\n    ... on File {\n      content\n    }\n  }  \n}\n": types.GetNodeFromPathDocument,
    "\nquery GetTokensFromAuth0($token: String!) {\n  getTokensFromAuth0(token: $token) {\n    access\n    refresh\n  }\n}\n": types.GetTokensFromAuth0Document,
    "\nquery Me {\n  me {\n    email\n  }\n}\n": types.MeDocument,
    "\nmutation RefreshTokens($refresh: String!) {\n  refreshTokens(refresh: $refresh) {\n    access\n    refresh\n  }\n}\n": types.RefreshTokensDocument,
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
export function gql(source: "\nquery GetNodeFromPath($path: String!) {\n  getNodeFromPath(path: $path) {\n    id\n    name\n    parent {\n      id\n      name\n    }\n    ...on Folder {\n      children {\n        id\n        name\n      }\n    }\n    ... on File {\n      content\n    }\n  }  \n}\n"): (typeof documents)["\nquery GetNodeFromPath($path: String!) {\n  getNodeFromPath(path: $path) {\n    id\n    name\n    parent {\n      id\n      name\n    }\n    ...on Folder {\n      children {\n        id\n        name\n      }\n    }\n    ... on File {\n      content\n    }\n  }  \n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery GetTokensFromAuth0($token: String!) {\n  getTokensFromAuth0(token: $token) {\n    access\n    refresh\n  }\n}\n"): (typeof documents)["\nquery GetTokensFromAuth0($token: String!) {\n  getTokensFromAuth0(token: $token) {\n    access\n    refresh\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery Me {\n  me {\n    email\n  }\n}\n"): (typeof documents)["\nquery Me {\n  me {\n    email\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nmutation RefreshTokens($refresh: String!) {\n  refreshTokens(refresh: $refresh) {\n    access\n    refresh\n  }\n}\n"): (typeof documents)["\nmutation RefreshTokens($refresh: String!) {\n  refreshTokens(refresh: $refresh) {\n    access\n    refresh\n  }\n}\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;