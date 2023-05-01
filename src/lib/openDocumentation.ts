import * as vscode from "vscode";
import { toast, toastType } from "./toast";

export const openDocumentation = () => {
  vscode.commands.executeCommand(
    "vscode.open",
    vscode.Uri.parse("https://github.com/revett/atlas#documentation")
  );

  toast(toastType.Success, "Opening documentation");
};
