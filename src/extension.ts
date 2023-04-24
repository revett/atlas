import * as dayjs from "dayjs";
import * as vscode from "vscode";
import { openDocumentation } from "./commands/openDocumentation";
import { createScratchNote } from "./commands/createScratchNote";
import { createTopicNote } from "./commands/createTopicNote";
var advancedFormatPlugin = require("dayjs/plugin/advancedFormat");
var timezonePlugin = require("dayjs/plugin/timezone");

export function activate(context: vscode.ExtensionContext) {
  dayjs.extend(advancedFormatPlugin);
  dayjs.extend(timezonePlugin);

  context.subscriptions.push(
    createScratchNote(),
    createTopicNote(),
    openDocumentation()
  );
}

export function deactivate() {}
