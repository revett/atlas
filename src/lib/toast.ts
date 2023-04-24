import * as vscode from "vscode";

export enum toastType {
  Error = "error",
  Success = "success",
}

export const toast = (type: toastType, message: string) => {
  const s = `📖 Atlas: ${message}`;

  switch (type) {
    case toastType.Error:
      vscode.window.showErrorMessage(s);
      break;
    case toastType.Success:
      vscode.window.showInformationMessage(s);
      break;
  }
};
