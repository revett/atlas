import * as vscode from "vscode";
import * as fs from "fs";
import { toastType, toast } from "../lib/toast";

export const createNote = async (filename: string, content: string) => {
  // Check that there is an open folder/workspace in VS Code.
  const workspaceFolder = vscode.workspace.workspaceFolders?.[0];
  if (!workspaceFolder) {
    await toast(
      toastType.Error,
      "Unable to create file as no folder/workspace is open in VS Code."
    );

    return;
  }

  // Create the full path to the new note in the workspace.
  const noteWorkspacePath = vscode.Uri.joinPath(workspaceFolder.uri, filename);

  // Check if file exists, if so open it.
  fs.access(noteWorkspacePath.path, fs.constants.F_OK, async (err) => {
    if (!err) {
      await toast(toastType.Success, `Note exists, opening "${filename}"`);

      const note = await vscode.workspace.openTextDocument(noteWorkspacePath);
      await vscode.window.showTextDocument(note);

      return;
    }

    // Create the new file, and insert the content.
    const editActions = new vscode.WorkspaceEdit();
    editActions.createFile(noteWorkspacePath, { ignoreIfExists: true });
    editActions.insert(noteWorkspacePath, new vscode.Position(0, 0), content);
    await vscode.workspace.applyEdit(editActions);

    await toast(toastType.Success, `Created "${filename}"`);

    // Save the new file.
    const note = await vscode.workspace.openTextDocument(noteWorkspacePath);
    await note.save();

    // Open the new file, and place the cursor on the last useful line.
    await vscode.window.showTextDocument(note);

    const lastLine = note.lineAt(note.lineCount - 2);
    const editor = vscode.window.activeTextEditor;

    if (editor) {
      const position = new vscode.Position(
        lastLine.range.end.line,
        lastLine.range.end.character
      );

      editor.selection = new vscode.Selection(position, position);
    }
  });
};
