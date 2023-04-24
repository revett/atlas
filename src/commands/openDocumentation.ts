import * as vscode from "vscode";
import { toastType, toast } from "../lib/toast";

export const openDocumentation = (): vscode.Disposable => {
  return vscode.commands.registerCommand("atlas.openDocumentation", () => {
    vscode.commands.executeCommand(
      "vscode.open",
      vscode.Uri.parse("https://github.com/revett/atlas#documentation")
    );

    toast(toastType.Success, "Opening documentation");
  });
};
