import * as vscode from "vscode";
import { toast, ToastType } from "./toast";

export const openDocumentation = () => {
  vscode.commands.executeCommand(
    "vscode.open",
    vscode.Uri.parse("https://github.com/revett/atlas#documentation")
  );

  toast(ToastType.success, "Opening documentation");
};
