package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewTeamHook(data *data.Data, logger log.Logger) repository.TeamHook {
	return &teamHookImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.team_hook")),
	}
}

type teamHookImpl struct {
	*data.Data
	helper *log.Helper
}

func (t *teamHookImpl) Create(ctx context.Context, hook do.NoticeHook) error {
	noticeHook := &team.NoticeHook{
		Name:    hook.GetName(),
		Remark:  hook.GetRemark(),
		Status:  hook.GetStatus(),
		URL:     hook.GetURL(),
		Method:  hook.GetMethod(),
		Secret:  hook.GetSecret(),
		Headers: hook.GetHeaders(),
		APP:     hook.GetApp(),
	}
	noticeHook.WithContext(ctx)

	query, _, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}

	return query.NoticeHook.WithContext(ctx).Create(noticeHook)
}

func (t *teamHookImpl) Update(ctx context.Context, hook do.NoticeHook) error {
	query, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	wrapper := []gen.Condition{
		query.NoticeHook.ID.Eq(hook.GetHookID()),
		query.NoticeHook.TeamID.Eq(teamID),
	}

	hookQuery := query.NoticeHook
	mutations := []field.AssignExpr{
		hookQuery.Name.Value(hook.GetName()),
		hookQuery.Remark.Value(hook.GetRemark()),
		hookQuery.Status.Value(hook.GetStatus().GetValue()),
		hookQuery.URL.Value(hook.GetURL()),
		hookQuery.Method.Value(hook.GetMethod().GetValue()),
		hookQuery.Secret.Value(hook.GetSecret()),
		hookQuery.Headers.Value(hook.GetHeaders()),
		hookQuery.APP.Value(hook.GetApp().GetValue()),
	}
	_, err = hookQuery.WithContext(ctx).Where(wrapper...).UpdateSimple(mutations...)
	return err
}

func (t *teamHookImpl) UpdateStatus(ctx context.Context, req *bo.UpdateTeamNoticeHookStatusRequest) error {
	query, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}

	wrapper := []gen.Condition{
		query.NoticeHook.ID.Eq(req.HookID),
		query.NoticeHook.TeamID.Eq(teamID),
	}

	hookQuery := query.NoticeHook
	_, err = hookQuery.WithContext(ctx).
		Where(wrapper...).
		UpdateSimple(hookQuery.Status.Value(req.Status.GetValue()))
	return err
}

func (t *teamHookImpl) Delete(ctx context.Context, hookID uint32) error {
	query, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}

	wrapper := []gen.Condition{
		query.NoticeHook.ID.Eq(hookID),
		query.NoticeHook.TeamID.Eq(teamID),
	}

	hookQuery := query.NoticeHook
	_, err = hookQuery.WithContext(ctx).Where(wrapper...).Delete()
	return err
}

func (t *teamHookImpl) Get(ctx context.Context, hookID uint32) (do.NoticeHook, error) {
	query, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}

	hookQuery := query.NoticeHook
	wrapper := []gen.Condition{
		hookQuery.ID.Eq(hookID),
		hookQuery.TeamID.Eq(teamID),
	}

	return hookQuery.WithContext(ctx).Where(wrapper...).First()
}

func (t *teamHookImpl) List(ctx context.Context, req *bo.ListTeamNoticeHookRequest) (*bo.ListTeamNoticeHookReply, error) {
	query, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}

	hookQuery := query.NoticeHook

	wrapper := hookQuery.WithContext(ctx).Where(hookQuery.TeamID.Eq(teamID))
	// Build conditions
	conditions := make([]gen.Condition, 0)
	if !req.Status.IsUnknown() {
		conditions = append(conditions, hookQuery.Status.Eq(req.Status.GetValue()))
	}
	if len(req.Apps) > 0 {
		conditions = append(conditions, hookQuery.APP.In(slices.Map(req.Apps, func(app vobj.HookApp) int8 { return app.GetValue() })...))
	}
	if !validate.TextIsNull(req.Keyword) {
		conditions = append(conditions, hookQuery.Name.Like(req.Keyword))
	}

	if validate.IsNotNil(req.PaginationRequest) {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		wrapper = wrapper.Offset(req.Offset()).Limit(int(req.Limit))
		req.WithTotal(total)
	}

	noticeHooks, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	hooks := slices.Map(noticeHooks, func(hook *team.NoticeHook) do.NoticeHook { return hook })
	return req.ToListTeamNoticeHookReply(hooks), nil
}
