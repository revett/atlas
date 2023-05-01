import * as vscode from "vscode";

export enum ToastType {
  error = "error",
  success = "success",
}

export const toast = (type: ToastType, message: string) => {
  const s = `📖 Atlas: ${message}`;

  switch (type) {
    case ToastType.error:
      vscode.window.showErrorMessage(s);
      break;
    case ToastType.success:
      vscode.window.showInformationMessage(s);
      break;
  }
};
