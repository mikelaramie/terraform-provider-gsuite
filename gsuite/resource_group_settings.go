package gsuite

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	groupSettings "google.golang.org/api/groupssettings/v1"
)

func resourceGroupSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupSettingsCreate,
		Read:   resourceGroupSettingsRead,
		Update: resourceGroupSettingsUpdate,
		Delete: resourceGroupSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGroupSettingsImporter,
		},

		Schema: map[string]*schema.Schema{
			"is_archived": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"allow_external_members": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"allow_google_communication": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"allow_web_posting": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "true",
			},
			"archive_only": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"custom_footer_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_reply_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"favorite_replies_on_top": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "true",
			},
			"include_custom_footer": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"include_in_global_address_list": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "true",
			},
			"max_message_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  26214400,
			},
			"members_can_post_as_the_group": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"message_display_font": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"DEFAULT_FONT", "FIXED_WIDTH_FONT", ""}, false),
				Default:      "DEFAULT_FONT",
			},
			"message_moderation_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"MODERATE_ALL_MESSAGES", "MODERATE_NON_MEMBERS", "MODERATE_NEW_MEMBERS", "MODERATE_NONE", ""}, false),
				Default:      "MODERATE_NONE",
			},
			"primary_language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reply_to": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"REPLY_TO_CUSTOM", "REPLY_TO_SENDER", "REPLY_TO_LIST", "REPLY_TO_OWNER", "REPLY_TO_IGNORE", "REPLY_TO_MANAGERS", ""}, false),
				Default:      "REPLY_TO_IGNORE",
			},
			"send_message_deny_notification": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"show_in_group_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"spam_moderation_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALLOW", "MODERATE", "SILENTLY_MODERATE", "REJECT", ""}, false),
				Default:      "MODERATE",
			},
			"who_can_add": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL_MANAGERS_CAN_ADD", "ALL_OWNERS_CAN_ADD", "ALL_MEMBERS_CAN_ADD", "NONE_CAN_ADD", ""}, false),
				Default:      "ALL_MANAGERS_CAN_ADD",
			},
			"who_can_add_references": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_approve_members": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL_OWNERS_CAN_APPROVE", "ALL_MANAGERS_CAN_APPROVE", "ALL_MEMBERS_CAN_APPROVE", "NONE_CAN_APPROVE", ""}, false),
				Default:      "ALL_MANAGERS_CAN_APPROVE",
			},
			"who_can_approve_messages": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_assign_topics": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_assist_content": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_ban_users": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_contact_owner": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ANYONE_CAN_CONTACT", "ALL_IN_DOMAIN_CAN_CONTACT", "ALL_MEMBERS_CAN_CONTACT", "ALL_MANAGERS_CAN_CONTACT", ""}, false),
				Default:      "ANYONE_CAN_CONTACT",
			},
			"who_can_delete_any_post": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_delete_topics": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_discover_group": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL_MEMBERS_CAN_DISCOVER", "ALL_IN_DOMAIN_CAN_DISCOVER", "ANYONE_CAN_DISCOVER", ""}, false),
				Default:      "ALL_MEMBERS_CAN_DISCOVER",
			},
			"who_can_enter_free_form_tags": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_hide_abuse": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_invite": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL_MEMBERS_CAN_INVITE", "ALL_MANAGERS_CAN_INVITE", "ALL_OWNERS_CAN_INVITE", "NONE_CAN_INVITE", ""}, false),
				Default:      "ALL_MANAGERS_CAN_INVITE",
			},
			"who_can_join": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ANYONE_CAN_JOIN", "ALL_IN_DOMAIN_CAN_JOIN", "INVITED_CAN_JOIN", "CAN_REQUEST_TO_JOIN", ""}, false),
				Default:      "CAN_REQUEST_TO_JOIN",
			},
			"who_can_leave_group": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL_MANAGERS_CAN_LEAVE", "ALL_OWNERS_CAN_LEAVE", "ALL_MEMBERS_CAN_LEAVE", "NONE_CAN_LEAVE", ""}, false),
				Default:      "ALL_MEMBERS_CAN_LEAVE",
			},
			"who_can_lock_topics": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_make_topics_sticky": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_mark_duplicate": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_mark_favorite_reply_on_any_topic": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_mark_favorite_reply_on_own_topic": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_mark_no_response_needed": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_moderate_content": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_moderate_members": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_modify_members": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_modify_tags_and_categories": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_move_topics_in": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_move_topics_out": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_post_announcements": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "OWNERS_AND_MANAGERS",
			},
			"who_can_post_message": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE_CAN_POST", "ALL_MANAGERS_CAN_POST", "ALL_MEMBERS_CAN_POST", "ALL_OWNERS_CAN_POST", "ALL_IN_DOMAIN_CAN_POST", "ANYONE_CAN_POST", ""}, false),
				Default:      "ANYONE_CAN_POST",
			},
			"who_can_take_topics": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_unassign_topic": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_unmark_favorite_reply_on_any_topic": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OWNERS_ONLY", "MANAGERS_ONLY", "OWNERS_AND_MANAGERS", "ALL_MEMBERS", ""}, false),
				Default:      "NONE",
			},
			"who_can_view_group": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ANYONE_CAN_VIEW", "ALL_IN_DOMAIN_CAN_VIEW", "ALL_MEMBERS_CAN_VIEW", "ALL_MANAGERS_CAN_VIEW", "ALL_OWNERS_CAN_VIEW", ""}, false),
				Default:      "ALL_MEMBERS_CAN_VIEW",
			},
			"who_can_view_membership": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL_IN_DOMAIN_CAN_VIEW", "ALL_MEMBERS_CAN_VIEW", "ALL_MANAGERS_CAN_VIEW", "ALL_OWNERS_CAN_VIEW", ""}, false),
				Default:      "ALL_MEMBERS_CAN_VIEW",
			},
		},
	}
}

func resourceGroupSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	// GroupSettings
	groupSetting := &groupSettings.Groups{
		Email: strings.ToLower(d.Get("email").(string)),
	}
	if v, ok := d.GetOk("allow_external_members"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "allow_external_members", v.(string))
		groupSetting.AllowExternalMembers = v.(string)
	}
	if v, ok := d.GetOk("allow_google_communication"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "allow_google_communication", v.(string))
		groupSetting.AllowGoogleCommunication = v.(string)
	}
	if v, ok := d.GetOk("allow_web_posting"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "allow_web_posting", v.(string))
		groupSetting.AllowWebPosting = v.(string)
	}
	if v, ok := d.GetOk("archive_only"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "archive_only", v.(string))
		groupSetting.ArchiveOnly = v.(string)
	}
	if v, ok := d.GetOk("custom_footer_text"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "custom_footer_text", v.(string))
		groupSetting.CustomFooterText = v.(string)
	}
	if v, ok := d.GetOk("custom_reply_to"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "custom_reply_to", v.(string))
		groupSetting.CustomReplyTo = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "description", v.(string))
		groupSetting.Description = v.(string)
	}
	if v, ok := d.GetOk("favorite_replies_on_top"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "favorite_replies_on_top", v.(string))
		groupSetting.FavoriteRepliesOnTop = v.(string)
	}
	if v, ok := d.GetOk("include_custom_footer"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "include_custom_footer", v.(string))
		groupSetting.IncludeCustomFooter = v.(string)
	}
	if v, ok := d.GetOk("include_in_global_address_list"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "include_in_global_address_list", v.(string))
		groupSetting.IncludeInGlobalAddressList = v.(string)
	}
	if v, ok := d.GetOk("max_message_bytes"); ok {
		log.Printf("[DEBUG] Setting %s: %d", "max_message_bytes", v.(int))
		groupSetting.MaxMessageBytes = int64(v.(int))
	}
	if v, ok := d.GetOk("members_can_post_as_the_group"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "members_can_post_as_the_group", v.(string))
		groupSetting.MembersCanPostAsTheGroup = v.(string)
	}
	if v, ok := d.GetOk("message_display_font"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "message_display_font", v.(string))
		groupSetting.MessageDisplayFont = v.(string)
	}
	if v, ok := d.GetOk("message_moderation_level"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "message_moderation_level", v.(string))
		groupSetting.MessageModerationLevel = v.(string)
	}
	if v, ok := d.GetOk("primary_language"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "primary_language", v.(string))
		groupSetting.PrimaryLanguage = v.(string)
	}
	if v, ok := d.GetOk("reply_to"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "reply_to", v.(string))
		groupSetting.ReplyTo = v.(string)
	}
	if v, ok := d.GetOk("send_message_deny_notification"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "send_message_deny_notification", v.(string))
		groupSetting.SendMessageDenyNotification = v.(string)
	}
	if v, ok := d.GetOk("show_in_group_directory"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "show_in_group_directory", v.(string))
		groupSetting.ShowInGroupDirectory = v.(string)
	}
	if v, ok := d.GetOk("spam_moderation_level"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "spam_moderation_level", v.(string))
		groupSetting.SpamModerationLevel = v.(string)
	}
	if v, ok := d.GetOk("who_can_add"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_add", v.(string))
		groupSetting.WhoCanAdd = v.(string)
	}
	if v, ok := d.GetOk("who_can_add_references"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_add_references", v.(string))
		groupSetting.WhoCanAddReferences = v.(string)
	}
	if v, ok := d.GetOk("who_can_approve_members"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_approve_members", v.(string))
		groupSetting.WhoCanApproveMembers = v.(string)
	}
	if v, ok := d.GetOk("who_can_approve_messages"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_approve_messages", v.(string))
		groupSetting.WhoCanApproveMessages = v.(string)
	}
	if v, ok := d.GetOk("who_can_assign_topics"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_assign_topics", v.(string))
		groupSetting.WhoCanAssignTopics = v.(string)
	}
	if v, ok := d.GetOk("who_can_assist_content"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_assist_content", v.(string))
		groupSetting.WhoCanAssistContent = v.(string)
	}
	if v, ok := d.GetOk("who_can_ban_users"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_ban_users", v.(string))
		groupSetting.WhoCanBanUsers = v.(string)
	}
	if v, ok := d.GetOk("who_can_contact_owner"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_contact_owner", v.(string))
		groupSetting.WhoCanContactOwner = v.(string)
	}
	if v, ok := d.GetOk("who_can_delete_any_post"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_delete_any_post", v.(string))
		groupSetting.WhoCanDeleteAnyPost = v.(string)
	}
	if v, ok := d.GetOk("who_can_delete_topics"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_delete_topics", v.(string))
		groupSetting.WhoCanDeleteTopics = v.(string)
	}
	if v, ok := d.GetOk("who_can_discover_group"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_discover_group", v.(string))
		groupSetting.WhoCanDiscoverGroup = v.(string)
	}
	if v, ok := d.GetOk("who_can_enter_free_form_tags"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_enter_free_form_tags", v.(string))
		groupSetting.WhoCanEnterFreeFormTags = v.(string)
	}
	if v, ok := d.GetOk("who_can_hide_abuse"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_hide_abuse", v.(string))
		groupSetting.WhoCanHideAbuse = v.(string)
	}
	if v, ok := d.GetOk("who_can_invite"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_invite", v.(string))
		groupSetting.WhoCanInvite = v.(string)
	}
	if v, ok := d.GetOk("who_can_join"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_join", v.(string))
		groupSetting.WhoCanJoin = v.(string)
	}
	if v, ok := d.GetOk("who_can_leave_group"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_leave_group", v.(string))
		groupSetting.WhoCanLeaveGroup = v.(string)
	}
	if v, ok := d.GetOk("who_can_lock_topics"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_lock_topics", v.(string))
		groupSetting.WhoCanLockTopics = v.(string)
	}
	if v, ok := d.GetOk("who_can_make_topics_sticky"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_make_topics_sticky", v.(string))
		groupSetting.WhoCanMakeTopicsSticky = v.(string)
	}
	if v, ok := d.GetOk("who_can_mark_duplicate"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_mark_duplicate", v.(string))
		groupSetting.WhoCanMarkDuplicate = v.(string)
	}
	if v, ok := d.GetOk("who_can_mark_favorite_reply_on_any_topic"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_mark_favorite_reply_on_any_topic", v.(string))
		groupSetting.WhoCanMarkFavoriteReplyOnAnyTopic = v.(string)
	}
	if v, ok := d.GetOk("who_can_mark_favorite_reply_on_own_topic"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_mark_favorite_reply_on_own_topic", v.(string))
		groupSetting.WhoCanMarkFavoriteReplyOnOwnTopic = v.(string)
	}
	if v, ok := d.GetOk("who_can_mark_no_response_needed"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_mark_no_response_needed", v.(string))
		groupSetting.WhoCanMarkNoResponseNeeded = v.(string)
	}
	if v, ok := d.GetOk("who_can_moderate_content"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_moderate_content", v.(string))
		groupSetting.WhoCanModerateContent = v.(string)
	}
	if v, ok := d.GetOk("who_can_moderate_members"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_moderate_members", v.(string))
		groupSetting.WhoCanModerateMembers = v.(string)
	}
	if v, ok := d.GetOk("who_can_modify_members"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_modify_members", v.(string))
		groupSetting.WhoCanModifyMembers = v.(string)
	}
	if v, ok := d.GetOk("who_can_modify_tags_and_categories"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_modify_tags_and_categories", v.(string))
		groupSetting.WhoCanModifyTagsAndCategories = v.(string)
	}
	if v, ok := d.GetOk("who_can_move_topics_in"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_move_topics_in", v.(string))
		groupSetting.WhoCanMoveTopicsIn = v.(string)
	}
	if v, ok := d.GetOk("who_can_move_topics_out"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_move_topics_out", v.(string))
		groupSetting.WhoCanMoveTopicsOut = v.(string)
	}
	if v, ok := d.GetOk("who_can_post_announcements"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_post_announcements", v.(string))
		groupSetting.WhoCanPostAnnouncements = v.(string)
	}
	if v, ok := d.GetOk("who_can_post_message"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_post_message", v.(string))
		groupSetting.WhoCanPostMessage = v.(string)
	}
	if v, ok := d.GetOk("who_can_take_topics"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_take_topics", v.(string))
		groupSetting.WhoCanTakeTopics = v.(string)
	}
	if v, ok := d.GetOk("who_can_unassign_topic"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_unassign_topic", v.(string))
		groupSetting.WhoCanUnassignTopic = v.(string)
	}
	if v, ok := d.GetOk("who_can_unmark_favorite_reply_on_any_topic"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_unmark_favorite_reply_on_any_topic", v.(string))
		groupSetting.WhoCanUnmarkFavoriteReplyOnAnyTopic = v.(string)
	}
	if v, ok := d.GetOk("who_can_view_group"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_view_group", v.(string))
		groupSetting.WhoCanViewGroup = v.(string)
	}
	if v, ok := d.GetOk("who_can_view_membership"); ok {
		log.Printf("[DEBUG] Setting %s: %s", "who_can_view_membership", v.(string))
		groupSetting.WhoCanViewMembership = v.(string)
	}

	var err error
	err = retry(func() error {
		_, err = config.groupSettings.Groups.Update(d.Get("email").(string), groupSetting).Do()
		return err
	})
	if err != nil {
		return fmt.Errorf("[ERROR] Something went wrong while updating group settings: %s", err)
	}

	return resourceGroupSettingsRead(d, meta)
}

func resourceGroupSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	// GroupSettings
	nullFields := []string{}
	groupSetting := &groupSettings.Groups{
		Email: strings.ToLower(d.Get("email").(string)),
	}
	if d.HasChange("allow_external_members") {
		if v, ok := d.GetOk("allow_external_members"); ok {
			log.Printf("[DEBUG] Updating group allow external members: %s", v.(string))
			groupSetting.AllowExternalMembers = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting AllowExternalMembers")
			groupSetting.AllowExternalMembers = ""
			nullFields = append(nullFields, "AllowExternalMembers")
		}
	}
	if d.HasChange("allow_google_communication") {
		if v, ok := d.GetOk("allow_google_communication"); ok {
			log.Printf("[DEBUG] Updating group allow google communication: %s", v.(string))
			groupSetting.AllowGoogleCommunication = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting AllowGoogleCommunication")
			groupSetting.AllowGoogleCommunication = ""
			nullFields = append(nullFields, "AllowGoogleCommunication")
		}
	}
	if d.HasChange("allow_web_posting") {
		if v, ok := d.GetOk("allow_web_posting"); ok {
			log.Printf("[DEBUG] Updating allow_web_posting: %s", v.(string))
			groupSetting.AllowWebPosting = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting AllowWebPosting")
			groupSetting.AllowWebPosting = ""
			nullFields = append(nullFields, "AllowWebPosting")
		}
	}
	if d.HasChange("archive_only") {
		if v, ok := d.GetOk("archive_only"); ok {
			log.Printf("[DEBUG] Updating archive_only: %s", v.(string))
			groupSetting.ArchiveOnly = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting ArchiveOnly")
			groupSetting.ArchiveOnly = ""
			nullFields = append(nullFields, "ArchiveOnly")
		}
	}
	if d.HasChange("custom_footer_text") {
		if v, ok := d.GetOk("custom_footer_text"); ok {
			log.Printf("[DEBUG] Updating custom_footer_text: %s", v.(string))
			groupSetting.CustomFooterText = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting CustomFooterText")
			groupSetting.CustomFooterText = ""
			nullFields = append(nullFields, "CustomFooterText")
		}
	}
	if d.HasChange("custom_reply_to") {
		if v, ok := d.GetOk("custom_reply_to"); ok {
			log.Printf("[DEBUG] Updating custom_reply_to: %s", v.(string))
			groupSetting.CustomReplyTo = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting CustomReplyTo")
			groupSetting.CustomReplyTo = ""
			nullFields = append(nullFields, "CustomReplyTo")
		}
	}
	if d.HasChange("description") {
		if v, ok := d.GetOk("description"); ok {
			log.Printf("[DEBUG] Updating description: %s", v.(string))
			groupSetting.Description = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting Description")
			groupSetting.Description = ""
			nullFields = append(nullFields, "Description")
		}
	}
	if d.HasChange("favorite_replies_on_top") {
		if v, ok := d.GetOk("favorite_replies_on_top"); ok {
			log.Printf("[DEBUG] Updating favorite_replies_on_top: %s", v.(string))
			groupSetting.FavoriteRepliesOnTop = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting FavoriteRepliesOnTop")
			groupSetting.FavoriteRepliesOnTop = ""
			nullFields = append(nullFields, "FavoriteRepliesOnTop")
		}
	}
	if d.HasChange("include_custom_footer") {
		if v, ok := d.GetOk("include_custom_footer"); ok {
			log.Printf("[DEBUG] Updating include_custom_footer: %s", v.(string))
			groupSetting.IncludeCustomFooter = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting IncludeCustomFooter")
			groupSetting.IncludeCustomFooter = ""
			nullFields = append(nullFields, "IncludeCustomFooter")
		}
	}
	if d.HasChange("include_in_global_address_list") {
		if v, ok := d.GetOk("include_in_global_address_list"); ok {
			log.Printf("[DEBUG] Updating include_in_global_address_list: %s", v.(string))
			groupSetting.IncludeInGlobalAddressList = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting IncludeInGlobalAddressList")
			groupSetting.IncludeInGlobalAddressList = ""
			nullFields = append(nullFields, "IncludeInGlobalAddressList")
		}
	}
	if d.HasChange("max_message_bytes") {
		if v, ok := d.GetOk("max_message_bytes"); ok {
			log.Printf("[DEBUG] Updating max_message_bytes: %d", v.(int))
			groupSetting.MaxMessageBytes = int64(v.(int))
		} else {
			log.Printf("[DEBUG] Removing groupSetting MaxMessageBytes")
			groupSetting.MaxMessageBytes = 0
			nullFields = append(nullFields, "MaxMessageBytes")
		}
	}
	if d.HasChange("members_can_post_as_the_group") {
		if v, ok := d.GetOk("members_can_post_as_the_group"); ok {
			log.Printf("[DEBUG] Updating members_can_post_as_the_group: %s", v.(string))
			groupSetting.MembersCanPostAsTheGroup = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting MembersCanPostAsTheGroup")
			groupSetting.MembersCanPostAsTheGroup = ""
			nullFields = append(nullFields, "MembersCanPostAsTheGroup")
		}
	}
	if d.HasChange("message_display_font") {
		if v, ok := d.GetOk("message_display_font"); ok {
			log.Printf("[DEBUG] Updating message_display_font: %s", v.(string))
			groupSetting.MessageDisplayFont = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting MessageDisplayFont")
			groupSetting.MessageDisplayFont = ""
			nullFields = append(nullFields, "MessageDisplayFont")
		}
	}
	if d.HasChange("message_moderation_level") {
		if v, ok := d.GetOk("message_moderation_level"); ok {
			log.Printf("[DEBUG] Updating message_moderation_level: %s", v.(string))
			groupSetting.MessageModerationLevel = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting MessageModerationLevel")
			groupSetting.MessageModerationLevel = ""
			nullFields = append(nullFields, "MessageModerationLevel")
		}
	}
	if d.HasChange("PrimaryLanguage") {
		if v, ok := d.GetOk("PrimaryLanguage"); ok {
			log.Printf("[DEBUG] Updating PrimaryLanguage: %s", v.(string))
			groupSetting.PrimaryLanguage = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting PrimaryLanguage")
			groupSetting.PrimaryLanguage = ""
			nullFields = append(nullFields, "PrimaryLanguage")
		}
	}
	if d.HasChange("reply_to") {
		if v, ok := d.GetOk("reply_to"); ok {
			log.Printf("[DEBUG] Updating reply_to: %s", v.(string))
			groupSetting.ReplyTo = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting ReplyTo")
			groupSetting.ReplyTo = ""
			nullFields = append(nullFields, "ReplyTo")
		}
	}
	if d.HasChange("send_message_deny_notification") {
		if v, ok := d.GetOk("send_message_deny_notification"); ok {
			log.Printf("[DEBUG] Updating send_message_deny_notification: %s", v.(string))
			groupSetting.SendMessageDenyNotification = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting SendMessageDenyNotification")
			groupSetting.SendMessageDenyNotification = ""
			nullFields = append(nullFields, "SendMessageDenyNotification")
		}
	}
	if d.HasChange("show_in_group_directory") {
		if v, ok := d.GetOk("show_in_group_directory"); ok {
			log.Printf("[DEBUG] Updating show_in_group_directory: %s", v.(string))
			groupSetting.ShowInGroupDirectory = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting ShowInGroupDirectory")
			groupSetting.ShowInGroupDirectory = ""
			nullFields = append(nullFields, "ShowInGroupDirectory")
		}
	}
	if d.HasChange("spam_moderation_level") {
		if v, ok := d.GetOk("spam_moderation_level"); ok {
			log.Printf("[DEBUG] Updating spam_moderation_level: %s", v.(string))
			groupSetting.SpamModerationLevel = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting SpamModerationLevel")
			groupSetting.SpamModerationLevel = ""
			nullFields = append(nullFields, "SpamModerationLevel")
		}
	}
	if d.HasChange("who_can_add") {
		if v, ok := d.GetOk("who_can_add"); ok {
			log.Printf("[DEBUG] Updating who_can_add: %s", v.(string))
			groupSetting.WhoCanAdd = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanAdd")
			groupSetting.WhoCanAdd = ""
			nullFields = append(nullFields, "WhoCanAdd")
		}
	}
	if d.HasChange("who_can_add_references") {
		if v, ok := d.GetOk("who_can_add_references"); ok {
			log.Printf("[DEBUG] Updating who_can_add_references: %s", v.(string))
			groupSetting.WhoCanAddReferences = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanAddReferences")
			groupSetting.WhoCanAddReferences = ""
			nullFields = append(nullFields, "WhoCanAddReferences")
		}
	}
	if d.HasChange("who_can_approve_members") {
		if v, ok := d.GetOk("who_can_approve_members"); ok {
			log.Printf("[DEBUG] Updating who_can_approve_members: %s", v.(string))
			groupSetting.WhoCanApproveMembers = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanApproveMembers")
			groupSetting.WhoCanApproveMembers = ""
			nullFields = append(nullFields, "WhoCanApproveMembers")
		}
	}
	if d.HasChange("who_can_approve_message") {
		if v, ok := d.GetOk("who_can_approve_message"); ok {
			log.Printf("[DEBUG] Updating who_can_approve_message: %s", v.(string))
			groupSetting.WhoCanApproveMessages = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanApproveMessages")
			groupSetting.WhoCanApproveMessages = ""
			nullFields = append(nullFields, "WhoCanApproveMessages")
		}
	}
	if d.HasChange("who_can_assign_topics") {
		if v, ok := d.GetOk("who_can_assign_topics"); ok {
			log.Printf("[DEBUG] Updating who_can_assign_topics: %s", v.(string))
			groupSetting.WhoCanAssignTopics = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanAssignTopics")
			groupSetting.WhoCanAssignTopics = ""
			nullFields = append(nullFields, "WhoCanAssignTopics")
		}
	}
	if d.HasChange("who_can_assist_content") {
		if v, ok := d.GetOk("who_can_assist_content"); ok {
			log.Printf("[DEBUG] Updating who_can_assist_content: %s", v.(string))
			groupSetting.WhoCanAssistContent = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanAssistContent")
			groupSetting.WhoCanAssistContent = ""
			nullFields = append(nullFields, "WhoCanAssistContent")
		}
	}
	if d.HasChange("who_can_ban_users") {
		if v, ok := d.GetOk("who_can_ban_users"); ok {
			log.Printf("[DEBUG] Updating who_can_ban_users: %s", v.(string))
			groupSetting.WhoCanBanUsers = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanBanUsers")
			groupSetting.WhoCanBanUsers = ""
			nullFields = append(nullFields, "WhoCanBanUsers")
		}
	}
	if d.HasChange("who_can_contact_owner") {
		if v, ok := d.GetOk("who_can_contact_owner"); ok {
			log.Printf("[DEBUG] Updating who_can_contact_owner: %s", v.(string))
			groupSetting.WhoCanContactOwner = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanContactOwner")
			groupSetting.WhoCanContactOwner = ""
			nullFields = append(nullFields, "WhoCanContactOwner")
		}
	}
	if d.HasChange("who_can_delete_any_post") {
		if v, ok := d.GetOk("who_can_delete_any_post"); ok {
			log.Printf("[DEBUG] Updating who_can_delete_any_post: %s", v.(string))
			groupSetting.WhoCanDeleteAnyPost = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanDeleteAnyPost")
			groupSetting.WhoCanDeleteAnyPost = ""
			nullFields = append(nullFields, "WhoCanDeleteAnyPost")
		}
	}
	if d.HasChange("who_can_delete_topics") {
		if v, ok := d.GetOk("who_can_delete_topics"); ok {
			log.Printf("[DEBUG] Updating who_can_delete_topics: %s", v.(string))
			groupSetting.WhoCanDeleteTopics = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanDeleteTopics")
			groupSetting.WhoCanDeleteTopics = ""
			nullFields = append(nullFields, "WhoCanDeleteTopics")
		}
	}
	if d.HasChange("who_can_discover_group") {
		if v, ok := d.GetOk("who_can_discover_group"); ok {
			log.Printf("[DEBUG] Updating who_can_discover_group: %s", v.(string))
			groupSetting.WhoCanDiscoverGroup = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanDiscoverGroup")
			groupSetting.WhoCanDiscoverGroup = ""
			nullFields = append(nullFields, "WhoCanDiscoverGroup")
		}
	}
	if d.HasChange("who_can_enter_free_form_tags") {
		if v, ok := d.GetOk("who_can_enter_free_form_tags"); ok {
			log.Printf("[DEBUG] Updating who_can_enter_free_form_tags: %s", v.(string))
			groupSetting.WhoCanEnterFreeFormTags = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanEnterFreeFormTags")
			groupSetting.WhoCanEnterFreeFormTags = ""
			nullFields = append(nullFields, "WhoCanEnterFreeFormTags")
		}
	}
	if d.HasChange("who_can_hide_abuse") {
		if v, ok := d.GetOk("who_can_hide_abuse"); ok {
			log.Printf("[DEBUG] Updating who_can_hide_abuse: %s", v.(string))
			groupSetting.WhoCanHideAbuse = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanHideAbuse")
			groupSetting.WhoCanHideAbuse = ""
			nullFields = append(nullFields, "WhoCanHideAbuse")
		}
	}
	if d.HasChange("who_can_invite") {
		if v, ok := d.GetOk("who_can_invite"); ok {
			log.Printf("[DEBUG] Updating who_can_invite: %s", v.(string))
			groupSetting.WhoCanInvite = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanInvite")
			groupSetting.WhoCanInvite = ""
			nullFields = append(nullFields, "WhoCanInvite")
		}
	}
	if d.HasChange("who_can_join") {
		if v, ok := d.GetOk("who_can_join"); ok {
			log.Printf("[DEBUG] Updating who_can_join: %s", v.(string))
			groupSetting.WhoCanJoin = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanJoin")
			groupSetting.WhoCanJoin = ""
			nullFields = append(nullFields, "WhoCanJoin")
		}
	}
	if d.HasChange("who_can_leave_group") {
		if v, ok := d.GetOk("who_can_leave_group"); ok {
			log.Printf("[DEBUG] Updating who_can_leave_group: %s", v.(string))
			groupSetting.WhoCanLeaveGroup = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanLeaveGroup")
			groupSetting.WhoCanLeaveGroup = ""
			nullFields = append(nullFields, "WhoCanLeaveGroup")
		}
	}
	if d.HasChange("who_can_lock_topics") {
		if v, ok := d.GetOk("who_can_lock_topics"); ok {
			log.Printf("[DEBUG] Updating who_can_lock_topics: %s", v.(string))
			groupSetting.WhoCanLockTopics = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanLockTopics")
			groupSetting.WhoCanLockTopics = ""
			nullFields = append(nullFields, "WhoCanLockTopics")
		}
	}
	if d.HasChange("who_can_make_topics_sticky") {
		if v, ok := d.GetOk("who_can_make_topics_sticky"); ok {
			log.Printf("[DEBUG] Updating who_can_make_topics_sticky: %s", v.(string))
			groupSetting.WhoCanMakeTopicsSticky = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanMakeTopicsSticky")
			groupSetting.WhoCanMakeTopicsSticky = ""
			nullFields = append(nullFields, "WhoCanMakeTopicsSticky")
		}
	}
	if d.HasChange("who_can_mark_duplicate") {
		if v, ok := d.GetOk("who_can_mark_duplicate"); ok {
			log.Printf("[DEBUG] Updating who_can_mark_duplicate: %s", v.(string))
			groupSetting.WhoCanMarkDuplicate = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanMarkDuplicate")
			groupSetting.WhoCanMarkDuplicate = ""
			nullFields = append(nullFields, "WhoCanMarkDuplicate")
		}
	}
	if d.HasChange("who_can_mark_favorite_reply_on_any_topic") {
		if v, ok := d.GetOk("who_can_mark_favorite_reply_on_any_topic"); ok {
			log.Printf("[DEBUG] Updating who_can_mark_favorite_reply_on_any_topic: %s", v.(string))
			groupSetting.WhoCanMarkFavoriteReplyOnAnyTopic = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanMarkFavoriteReplyOnAnyTopic")
			groupSetting.WhoCanMarkFavoriteReplyOnAnyTopic = ""
			nullFields = append(nullFields, "WhoCanMarkFavoriteReplyOnAnyTopic")
		}
	}
	if d.HasChange("who_can_mark_favorite_reply_on_own_topic") {
		if v, ok := d.GetOk("who_can_mark_favorite_reply_on_own_topic"); ok {
			log.Printf("[DEBUG] Updating who_can_mark_favorite_reply_on_own_topic: %s", v.(string))
			groupSetting.WhoCanMarkFavoriteReplyOnOwnTopic = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanMarkFavoriteReplyOnOwnTopic")
			groupSetting.WhoCanMarkFavoriteReplyOnOwnTopic = ""
			nullFields = append(nullFields, "WhoCanMarkFavoriteReplyOnOwnTopic")
		}
	}
	if d.HasChange("who_can_mark_no_response_needed") {
		if v, ok := d.GetOk("who_can_mark_no_response_needed"); ok {
			log.Printf("[DEBUG] Updating who_can_mark_no_response_needed: %s", v.(string))
			groupSetting.WhoCanMarkNoResponseNeeded = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanMarkNoResponseNeeded")
			groupSetting.WhoCanMarkNoResponseNeeded = ""
			nullFields = append(nullFields, "WhoCanMarkNoResponseNeeded")
		}
	}
	if d.HasChange("who_can_moderate_content") {
		if v, ok := d.GetOk("who_can_moderate_content"); ok {
			log.Printf("[DEBUG] Updating who_can_moderate_content: %s", v.(string))
			groupSetting.WhoCanModerateContent = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanModerateContent")
			groupSetting.WhoCanModerateContent = ""
			nullFields = append(nullFields, "WhoCanModerateContent")
		}
	}
	if d.HasChange("who_can_moderate_members") {
		if v, ok := d.GetOk("who_can_moderate_members"); ok {
			log.Printf("[DEBUG] Updating who_can_moderate_members: %s", v.(string))
			groupSetting.WhoCanModerateMembers = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanModerateMembers")
			groupSetting.WhoCanModerateMembers = ""
			nullFields = append(nullFields, "WhoCanModerateMembers")
		}
	}
	if d.HasChange("who_can_modify_members") {
		if v, ok := d.GetOk("who_can_modify_members"); ok {
			log.Printf("[DEBUG] Updating who_can_modify_members: %s", v.(string))
			groupSetting.WhoCanModifyMembers = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanModifyMembers")
			groupSetting.WhoCanModifyMembers = ""
			nullFields = append(nullFields, "WhoCanModifyMembers")
		}
	}
	if d.HasChange("who_can_modify_tags_and_categories") {
		if v, ok := d.GetOk("who_can_modify_tags_and_categories"); ok {
			log.Printf("[DEBUG] Updating who_can_modify_tags_and_categories: %s", v.(string))
			groupSetting.WhoCanModifyTagsAndCategories = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanModifyTagsAndCategories")
			groupSetting.WhoCanModifyTagsAndCategories = ""
			nullFields = append(nullFields, "WhoCanModifyTagsAndCategories")
		}
	}
	if d.HasChange("who_can_move_topics_in") {
		if v, ok := d.GetOk("who_can_move_topics_in"); ok {
			log.Printf("[DEBUG] Updating who_can_move_topics_in: %s", v.(string))
			groupSetting.WhoCanMoveTopicsIn = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanMoveTopicsIn")
			groupSetting.WhoCanMoveTopicsIn = ""
			nullFields = append(nullFields, "WhoCanMoveTopicsIn")
		}
	}
	if d.HasChange("who_can_move_topics_out") {
		if v, ok := d.GetOk("who_can_move_topics_out"); ok {
			log.Printf("[DEBUG] Updating who_can_move_topics_out: %s", v.(string))
			groupSetting.WhoCanMoveTopicsOut = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanMoveTopicsOut")
			groupSetting.WhoCanMoveTopicsOut = ""
			nullFields = append(nullFields, "WhoCanMoveTopicsOut")
		}
	}
	if d.HasChange("who_can_post_announcements") {
		if v, ok := d.GetOk("who_can_post_announcements"); ok {
			log.Printf("[DEBUG] Updating who_can_post_announcements: %s", v.(string))
			groupSetting.WhoCanPostAnnouncements = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanPostAnnouncements")
			groupSetting.WhoCanPostAnnouncements = ""
			nullFields = append(nullFields, "WhoCanPostAnnouncements")
		}
	}
	if d.HasChange("who_can_post_message") {
		if v, ok := d.GetOk("who_can_post_message"); ok {
			log.Printf("[DEBUG] Updating who_can_post_message: %s", v.(string))
			groupSetting.WhoCanPostMessage = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanPostMessage")
			groupSetting.WhoCanPostMessage = ""
			nullFields = append(nullFields, "WhoCanPostMessage")
		}
	}
	if d.HasChange("who_can_take_topics") {
		if v, ok := d.GetOk("who_can_take_topics"); ok {
			log.Printf("[DEBUG] Updating who_can_take_topics: %s", v.(string))
			groupSetting.WhoCanTakeTopics = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanTakeTopics")
			groupSetting.WhoCanTakeTopics = ""
			nullFields = append(nullFields, "WhoCanTakeTopics")
		}
	}
	if d.HasChange("who_can_unassign_topic") {
		if v, ok := d.GetOk("who_can_unassign_topic"); ok {
			log.Printf("[DEBUG] Updating who_can_unassign_topic: %s", v.(string))
			groupSetting.WhoCanUnassignTopic = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanUnassignTopic")
			groupSetting.WhoCanUnassignTopic = ""
			nullFields = append(nullFields, "WhoCanUnassignTopic")
		}
	}
	if d.HasChange("who_can_unmark_favorite_reply_on_any_topic") {
		if v, ok := d.GetOk("who_can_unmark_favorite_reply_on_any_topic"); ok {
			log.Printf("[DEBUG] Updating who_can_unmark_favorite_reply_on_any_topic: %s", v.(string))
			groupSetting.WhoCanUnmarkFavoriteReplyOnAnyTopic = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanUnmarkFavoriteReplyOnAnyTopic")
			groupSetting.WhoCanUnmarkFavoriteReplyOnAnyTopic = ""
			nullFields = append(nullFields, "WhoCanUnmarkFavoriteReplyOnAnyTopic")
		}
	}
	if d.HasChange("who_can_view_group") {
		if v, ok := d.GetOk("who_can_view_group"); ok {
			log.Printf("[DEBUG] Updating who_can_view_group: %s", v.(string))
			groupSetting.WhoCanViewGroup = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanViewGroup")
			groupSetting.WhoCanViewGroup = ""
			nullFields = append(nullFields, "WhoCanViewGroup")
		}
	}
	if d.HasChange("who_can_view_membership") {
		if v, ok := d.GetOk("who_can_view_membership"); ok {
			log.Printf("[DEBUG] Updating who_can_view_membership: %s", v.(string))
			groupSetting.WhoCanViewMembership = v.(string)
		} else {
			log.Printf("[DEBUG] Removing groupSetting WhoCanViewMembership")
			groupSetting.WhoCanViewMembership = ""
			nullFields = append(nullFields, "WhoCanViewMembership")
		}
	}

	var err error
	err = retry(func() error {
		_, err = config.groupSettings.Groups.Update(d.Get("email").(string), groupSetting).Do()
		return err
	})

	if err != nil {
		return fmt.Errorf("[ERROR] Error updating group settings: %s", err)
	}

	return resourceGroupSettingsRead(d, meta)
}

func resourceGroupSettingsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var err error
	var groupSetting *groupSettings.Groups
	err = retry(func() error {
		groupSetting, err = config.groupSettings.Groups.Get(d.Get("email").(string)).Do()
		return err
	})

	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("Group Settings for %q", d.Get("name").(string)))
	}

	d.SetId(d.Get("email").(string))
	d.Set("allow_external_members", groupSetting.AllowExternalMembers)
	d.Set("allow_google_communication", groupSetting.AllowGoogleCommunication)
	d.Set("allow_web_posting", groupSetting.AllowWebPosting)
	d.Set("archive_only", groupSetting.ArchiveOnly)
	d.Set("custom_footer_text", groupSetting.CustomFooterText)
	d.Set("custom_reply_to", groupSetting.CustomReplyTo)
	d.Set("description", groupSetting.Description)
	d.Set("favorite_replies_on_top", groupSetting.FavoriteRepliesOnTop)
	d.Set("include_custom_footer", groupSetting.IncludeCustomFooter)
	d.Set("include_in_global_address_list", groupSetting.IncludeInGlobalAddressList)
	d.Set("max_message_bytes", groupSetting.MaxMessageBytes)
	d.Set("members_can_post_as_the_group", groupSetting.MembersCanPostAsTheGroup)
	d.Set("message_display_font", groupSetting.MessageDisplayFont)
	d.Set("message_moderation_level", groupSetting.MessageModerationLevel)
	d.Set("primary_language", groupSetting.PrimaryLanguage)
	d.Set("reply_to", groupSetting.ReplyTo)
	d.Set("send_message_deny_notification", groupSetting.SendMessageDenyNotification)
	d.Set("show_in_group_directory", groupSetting.ShowInGroupDirectory)
	d.Set("spam_moderation_level", groupSetting.SpamModerationLevel)
	d.Set("who_can_add", groupSetting.WhoCanAdd)
	d.Set("who_can_add_references", groupSetting.WhoCanAddReferences)
	d.Set("who_can_approve_members", groupSetting.WhoCanApproveMembers)
	d.Set("who_can_approve_messages", groupSetting.WhoCanApproveMessages)
	d.Set("who_can_assign_topics", groupSetting.WhoCanAssignTopics)
	d.Set("who_can_assist_content", groupSetting.WhoCanAssistContent)
	d.Set("who_can_ban_users", groupSetting.WhoCanBanUsers)
	d.Set("who_can_contact_owner", groupSetting.WhoCanContactOwner)
	d.Set("who_can_delete_any_post", groupSetting.WhoCanDeleteAnyPost)
	d.Set("who_can_delete_topics", groupSetting.WhoCanDeleteTopics)
	d.Set("who_can_discover_group", groupSetting.WhoCanDiscoverGroup)
	d.Set("who_can_enter_free_form_tags", groupSetting.WhoCanEnterFreeFormTags)
	d.Set("who_can_hide_abuse", groupSetting.WhoCanHideAbuse)
	d.Set("who_can_invite", groupSetting.WhoCanInvite)
	d.Set("who_can_join", groupSetting.WhoCanJoin)
	d.Set("who_can_leave_group", groupSetting.WhoCanLeaveGroup)
	d.Set("who_can_lock_topics", groupSetting.WhoCanLockTopics)
	d.Set("who_can_make_topics_sticky", groupSetting.WhoCanMakeTopicsSticky)
	d.Set("who_can_mark_duplicate", groupSetting.WhoCanMarkDuplicate)
	d.Set("who_can_mark_favorite_reply_on_any_topic", groupSetting.WhoCanMarkFavoriteReplyOnAnyTopic)
	d.Set("who_can_mark_favorite_reply_on_own_topic", groupSetting.WhoCanMarkFavoriteReplyOnOwnTopic)
	d.Set("who_can_mark_no_response_needed", groupSetting.WhoCanMarkNoResponseNeeded)
	d.Set("who_can_moderate_content", groupSetting.WhoCanModerateContent)
	d.Set("who_can_moderate_members", groupSetting.WhoCanModerateMembers)
	d.Set("who_can_modify_members", groupSetting.WhoCanModifyMembers)
	d.Set("who_can_modify_tags_and_categories", groupSetting.WhoCanModifyTagsAndCategories)
	d.Set("who_can_move_topics_in", groupSetting.WhoCanMoveTopicsIn)
	d.Set("who_can_move_topics_out", groupSetting.WhoCanMoveTopicsOut)
	d.Set("who_can_post_announcements", groupSetting.WhoCanPostAnnouncements)
	d.Set("who_can_post_message", groupSetting.WhoCanPostMessage)
	d.Set("who_can_take_topics", groupSetting.WhoCanTakeTopics)
	d.Set("who_can_unassign_topic", groupSetting.WhoCanUnassignTopic)
	d.Set("who_can_unmark_favorite_reply_on_any_topic", groupSetting.WhoCanUnmarkFavoriteReplyOnAnyTopic)
	d.Set("who_can_view_group", groupSetting.WhoCanViewGroup)
	d.Set("who_can_view_membership", groupSetting.WhoCanViewMembership)

	return nil
}

func resourceGroupSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var err error
	err = retry(func() error {
		err = config.directory.Groups.Delete(d.Id()).Do()
		return err
	})
	if err != nil {
		return fmt.Errorf("[ERROR] Error deleting group: %s", err)
	}

	d.SetId("")
	return nil
}

// Allow importing using email
func resourceGroupSettingsImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	id, err := config.groupSettings.Groups.Get(d.Id()).Do()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error fetching group settings. Make sure the group exists: %s ", err)
	}

	d.SetId(d.Id())
	d.Set("allow_external_members", id.AllowExternalMembers)
	d.Set("allow_google_communication", id.AllowGoogleCommunication)
	d.Set("allow_web_posting", id.AllowWebPosting)
	d.Set("archive_only", id.ArchiveOnly)
	d.Set("custom_footer_text", id.CustomFooterText)
	d.Set("custom_reply_to", id.CustomReplyTo)
	d.Set("description", id.Description)
	d.Set("email", id.Email)
	d.Set("favorite_replies_on_top", id.FavoriteRepliesOnTop)
	d.Set("include_custom_footer", id.IncludeCustomFooter)
	d.Set("include_in_global_address_list", id.IncludeInGlobalAddressList)
	d.Set("max_message_bytes", id.MaxMessageBytes)
	d.Set("members_can_post_as_the_group", id.MembersCanPostAsTheGroup)
	d.Set("message_display_font", id.MessageDisplayFont)
	d.Set("message_moderation_level", id.MessageModerationLevel)
	d.Set("primary_language", id.PrimaryLanguage)
	d.Set("reply_to", id.ReplyTo)
	d.Set("send_message_deny_notification", id.SendMessageDenyNotification)
	d.Set("show_in_group_directory", id.ShowInGroupDirectory)
	d.Set("spam_moderation_level", id.SpamModerationLevel)
	d.Set("who_can_add", id.WhoCanAdd)
	d.Set("who_can_add_references", id.WhoCanAddReferences)
	d.Set("who_can_approve_members", id.WhoCanApproveMembers)
	d.Set("who_can_approve_messages", id.WhoCanApproveMessages)
	d.Set("who_can_assign_topics", id.WhoCanAssignTopics)
	d.Set("who_can_assist_content", id.WhoCanAssistContent)
	d.Set("who_can_ban_users", id.WhoCanBanUsers)
	d.Set("who_can_contact_owner", id.WhoCanContactOwner)
	d.Set("who_can_delete_any_post", id.WhoCanDeleteAnyPost)
	d.Set("who_can_delete_topics", id.WhoCanDeleteTopics)
	d.Set("who_can_discover_group", id.WhoCanDiscoverGroup)
	d.Set("who_can_enter_free_form_tags", id.WhoCanEnterFreeFormTags)
	d.Set("who_can_hide_abuse", id.WhoCanHideAbuse)
	d.Set("who_can_invite", id.WhoCanInvite)
	d.Set("who_can_join", id.WhoCanJoin)
	d.Set("who_can_leave_group", id.WhoCanLeaveGroup)
	d.Set("who_can_lock_topics", id.WhoCanLockTopics)
	d.Set("who_can_make_topics_sticky", id.WhoCanMakeTopicsSticky)
	d.Set("who_can_mark_duplicate", id.WhoCanMarkDuplicate)
	d.Set("who_can_mark_favorite_reply_on_any_topic", id.WhoCanMarkFavoriteReplyOnAnyTopic)
	d.Set("who_can_mark_favorite_reply_on_own_topic", id.WhoCanMarkFavoriteReplyOnOwnTopic)
	d.Set("who_can_mark_no_response_needed", id.WhoCanMarkNoResponseNeeded)
	d.Set("who_can_moderate_content", id.WhoCanModerateContent)
	d.Set("who_can_moderate_members", id.WhoCanModerateMembers)
	d.Set("who_can_modify_members", id.WhoCanModifyMembers)
	d.Set("who_can_modify_tags_and_categories", id.WhoCanModifyTagsAndCategories)
	d.Set("who_can_move_topics_in", id.WhoCanMoveTopicsIn)
	d.Set("who_can_move_topics_out", id.WhoCanMoveTopicsOut)
	d.Set("who_can_post_announcements", id.WhoCanPostAnnouncements)
	d.Set("who_can_post_message", id.WhoCanPostMessage)
	d.Set("who_can_take_topics", id.WhoCanTakeTopics)
	d.Set("who_can_unassign_topic", id.WhoCanUnassignTopic)
	d.Set("who_can_unmark_favorite_reply_on_any_topic", id.WhoCanUnmarkFavoriteReplyOnAnyTopic)
	d.Set("who_can_view_group", id.WhoCanViewGroup)
	d.Set("who_can_view_membership", id.WhoCanViewMembership)

	return []*schema.ResourceData{d}, nil
}
