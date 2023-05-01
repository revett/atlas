import * as vscode from "vscode";
import * as openDocumentationLib from "../lib/openDocumentation";

export const openDocumentation = (): vscode.Disposable => {
  return vscode.commands.registerCommand("atlas.openDocumentation", () => {
    openDocumentationLib.openDocumentation();
  });
};
