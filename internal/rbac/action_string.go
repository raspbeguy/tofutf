// Code generated by "stringer -type Action ./internal/rbac"; DO NOT EDIT.

package rbac

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WatchAction-0]
	_ = x[CreateOrganizationAction-1]
	_ = x[UpdateOrganizationAction-2]
	_ = x[GetOrganizationAction-3]
	_ = x[ListOrganizationsAction-4]
	_ = x[GetEntitlementsAction-5]
	_ = x[DeleteOrganizationAction-6]
	_ = x[CreateVCSProviderAction-7]
	_ = x[GetVCSProviderAction-8]
	_ = x[ListVCSProvidersAction-9]
	_ = x[DeleteVCSProviderAction-10]
	_ = x[CreateAgentTokenAction-11]
	_ = x[ListAgentTokensAction-12]
	_ = x[DeleteAgentTokenAction-13]
	_ = x[CreateOrganizationTokenAction-14]
	_ = x[DeleteOrganizationTokenAction-15]
	_ = x[CreateRunTokenAction-16]
	_ = x[CreateModuleAction-17]
	_ = x[CreateModuleVersionAction-18]
	_ = x[UpdateModuleAction-19]
	_ = x[ListModulesAction-20]
	_ = x[GetModuleAction-21]
	_ = x[DeleteModuleAction-22]
	_ = x[DeleteModuleVersionAction-23]
	_ = x[CreateWorkspaceVariableAction-24]
	_ = x[UpdateWorkspaceVariableAction-25]
	_ = x[ListWorkspaceVariablesAction-26]
	_ = x[GetWorkspaceVariableAction-27]
	_ = x[DeleteWorkspaceVariableAction-28]
	_ = x[CreateVariableSetAction-29]
	_ = x[UpdateVariableSetAction-30]
	_ = x[ListVariableSetsAction-31]
	_ = x[GetVariableSetAction-32]
	_ = x[DeleteVariableSetAction-33]
	_ = x[CreateVariableSetVariableAction-34]
	_ = x[UpdateVariableSetVariableAction-35]
	_ = x[GetVariableSetVariableAction-36]
	_ = x[DeleteVariableSetVariableAction-37]
	_ = x[AddVariableToSetAction-38]
	_ = x[RemoveVariableFromSetAction-39]
	_ = x[ApplyVariableSetToWorkspacesAction-40]
	_ = x[DeleteVariableSetFromWorkspacesAction-41]
	_ = x[GetRunAction-42]
	_ = x[ListRunsAction-43]
	_ = x[ApplyRunAction-44]
	_ = x[CreateRunAction-45]
	_ = x[DiscardRunAction-46]
	_ = x[DeleteRunAction-47]
	_ = x[CancelRunAction-48]
	_ = x[EnqueuePlanAction-49]
	_ = x[StartPhaseAction-50]
	_ = x[FinishPhaseAction-51]
	_ = x[PutChunkAction-52]
	_ = x[TailLogsAction-53]
	_ = x[GetPlanFileAction-54]
	_ = x[UploadPlanFileAction-55]
	_ = x[GetLockFileAction-56]
	_ = x[UploadLockFileAction-57]
	_ = x[ListWorkspacesAction-58]
	_ = x[GetWorkspaceAction-59]
	_ = x[CreateWorkspaceAction-60]
	_ = x[DeleteWorkspaceAction-61]
	_ = x[SetWorkspacePermissionAction-62]
	_ = x[UnsetWorkspacePermissionAction-63]
	_ = x[UpdateWorkspaceAction-64]
	_ = x[ListTagsAction-65]
	_ = x[DeleteTagsAction-66]
	_ = x[TagWorkspacesAction-67]
	_ = x[AddTagsAction-68]
	_ = x[RemoveTagsAction-69]
	_ = x[ListWorkspaceTags-70]
	_ = x[LockWorkspaceAction-71]
	_ = x[UnlockWorkspaceAction-72]
	_ = x[ForceUnlockWorkspaceAction-73]
	_ = x[CreateStateVersionAction-74]
	_ = x[ListStateVersionsAction-75]
	_ = x[GetStateVersionAction-76]
	_ = x[DeleteStateVersionAction-77]
	_ = x[RollbackStateVersionAction-78]
	_ = x[DownloadStateAction-79]
	_ = x[GetStateVersionOutputAction-80]
	_ = x[CreateConfigurationVersionAction-81]
	_ = x[ListConfigurationVersionsAction-82]
	_ = x[GetConfigurationVersionAction-83]
	_ = x[DownloadConfigurationVersionAction-84]
	_ = x[DeleteConfigurationVersionAction-85]
	_ = x[CreateUserAction-86]
	_ = x[ListUsersAction-87]
	_ = x[GetUserAction-88]
	_ = x[DeleteUserAction-89]
	_ = x[CreateTeamAction-90]
	_ = x[UpdateTeamAction-91]
	_ = x[GetTeamAction-92]
	_ = x[ListTeamsAction-93]
	_ = x[DeleteTeamAction-94]
	_ = x[AddTeamMembershipAction-95]
	_ = x[RemoveTeamMembershipAction-96]
	_ = x[CreateNotificationConfigurationAction-97]
	_ = x[UpdateNotificationConfigurationAction-98]
	_ = x[ListNotificationConfigurationsAction-99]
	_ = x[GetNotificationConfigurationAction-100]
	_ = x[DeleteNotificationConfigurationAction-101]
	_ = x[CreateGithubAppAction-102]
	_ = x[UpdateGithubAppAction-103]
	_ = x[GetGithubAppAction-104]
	_ = x[ListGithubAppsAction-105]
	_ = x[DeleteGithubAppAction-106]
	_ = x[CreateGithubAppInstallAction-107]
	_ = x[DeleteGithubAppInstallAction-108]
}

const _Action_name = "WatchActionCreateOrganizationActionUpdateOrganizationActionGetOrganizationActionListOrganizationsActionGetEntitlementsActionDeleteOrganizationActionCreateVCSProviderActionGetVCSProviderActionListVCSProvidersActionDeleteVCSProviderActionCreateAgentTokenActionListAgentTokensActionDeleteAgentTokenActionCreateOrganizationTokenActionDeleteOrganizationTokenActionCreateRunTokenActionCreateModuleActionCreateModuleVersionActionUpdateModuleActionListModulesActionGetModuleActionDeleteModuleActionDeleteModuleVersionActionCreateWorkspaceVariableActionUpdateWorkspaceVariableActionListWorkspaceVariablesActionGetWorkspaceVariableActionDeleteWorkspaceVariableActionCreateVariableSetActionUpdateVariableSetActionListVariableSetsActionGetVariableSetActionDeleteVariableSetActionCreateVariableSetVariableActionUpdateVariableSetVariableActionGetVariableSetVariableActionDeleteVariableSetVariableActionAddVariableToSetActionRemoveVariableFromSetActionApplyVariableSetToWorkspacesActionDeleteVariableSetFromWorkspacesActionGetRunActionListRunsActionApplyRunActionCreateRunActionDiscardRunActionDeleteRunActionCancelRunActionEnqueuePlanActionStartPhaseActionFinishPhaseActionPutChunkActionTailLogsActionGetPlanFileActionUploadPlanFileActionGetLockFileActionUploadLockFileActionListWorkspacesActionGetWorkspaceActionCreateWorkspaceActionDeleteWorkspaceActionSetWorkspacePermissionActionUnsetWorkspacePermissionActionUpdateWorkspaceActionListTagsActionDeleteTagsActionTagWorkspacesActionAddTagsActionRemoveTagsActionListWorkspaceTagsLockWorkspaceActionUnlockWorkspaceActionForceUnlockWorkspaceActionCreateStateVersionActionListStateVersionsActionGetStateVersionActionDeleteStateVersionActionRollbackStateVersionActionDownloadStateActionGetStateVersionOutputActionCreateConfigurationVersionActionListConfigurationVersionsActionGetConfigurationVersionActionDownloadConfigurationVersionActionDeleteConfigurationVersionActionCreateUserActionListUsersActionGetUserActionDeleteUserActionCreateTeamActionUpdateTeamActionGetTeamActionListTeamsActionDeleteTeamActionAddTeamMembershipActionRemoveTeamMembershipActionCreateNotificationConfigurationActionUpdateNotificationConfigurationActionListNotificationConfigurationsActionGetNotificationConfigurationActionDeleteNotificationConfigurationActionCreateGithubAppActionUpdateGithubAppActionGetGithubAppActionListGithubAppsActionDeleteGithubAppActionCreateGithubAppInstallActionDeleteGithubAppInstallAction"

var _Action_index = [...]uint16{0, 11, 35, 59, 80, 103, 124, 148, 171, 191, 213, 236, 258, 279, 301, 330, 359, 379, 397, 422, 440, 457, 472, 490, 515, 544, 573, 601, 627, 656, 679, 702, 724, 744, 767, 798, 829, 857, 888, 910, 937, 971, 1008, 1020, 1034, 1048, 1063, 1079, 1094, 1109, 1126, 1142, 1159, 1173, 1187, 1204, 1224, 1241, 1261, 1281, 1299, 1320, 1341, 1369, 1399, 1420, 1434, 1450, 1469, 1482, 1498, 1515, 1534, 1555, 1581, 1605, 1628, 1649, 1673, 1699, 1718, 1745, 1777, 1808, 1837, 1871, 1903, 1919, 1934, 1947, 1963, 1979, 1995, 2008, 2023, 2039, 2062, 2088, 2125, 2162, 2198, 2232, 2269, 2290, 2311, 2329, 2349, 2370, 2398, 2426}

func (i Action) String() string {
	if i < 0 || i >= Action(len(_Action_index)-1) {
		return "Action(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Action_name[_Action_index[i]:_Action_index[i+1]]
}
