package tests

import (
	"github.com/sunfmin/excerpt"
	"testing"
	"strings"
)

func TestHighlightHtml(t *testing.T) {
	var h = `
<p>For all interested: let’s start gathering our ideas about how we might integrate GitHub.</p>

<p>I imagine this to be an optional add-on for a Qortex installation that can be configured for each group with preferences to enable one or more of the features we support.</p>

<h3>Configuration</h3>

<p>Something like this:</p>

<p><strong>Enable GitHub Integration</strong></p>

<pre><code>API key: [ .... ]

[x] Enable Issue Tracking (perhaps with additional preferences)
[x] Pull GitHub Commits into this group
...
</code></pre>

<h3>Issues (Tickets)</h3>

<p><a href="http://qortex.net/theplant.jp/#groups/4fd78138558fbe76ff000028/entry/50823d763c58164a9500344c">As mentioned by Oli</a>, an integration for Issues could serve as a valuable bridge between developers and managers/designers as well as allowing an organization to consolidate knowledge around products in Qortex.</p>

<p>GitHub’s API seems to be fairly rich: <a href="http://developer.github.com/v3/issues/" title="Link: http://developer.github.com/v3/issues/" target="_blank">http://developer.github.com/v3/issues/</a></p>

<h3>First thoughts:</h3>

<ul>
<li>Keep the creation of milestones and labels in GitHub<br></li>
<li>Add a new “Action” option: “Create an Issue for this on GitHub”. If checked, include additional options to assign the issue to a person, set the tag, and milestone. This would not replace the Notification system, I think it would still be valuable to notify others even if the ticket is not assigned to them.<br></li>
<li>Sync issues and comments on issues from GitHub and Qortex.<br></li>
<li>Add a new “Issues” tab on Groups that have this option enabled.<br></li>
<li>Add a My Issues item in the left nav if any issues are assigned to you.<br></li>
</ul>

<h3>GitHub Commits</h3>

<p>I’m not so sure about this one apart from pulling commit messages into Qortex. Maybe there are other interesting things that can be done?</p>

`

	r, _, err := excerpt.HighlightHtml(h, []string{"github"}, highlight)
	if err != nil {
		t.Error(err)
	}
	if strings.Index(r, "http://developer.github.com/v3/issues/") < 0 {
		t.Error(r)
	}
	if strings.Index(r, "might integrate *GitHub*") < 0 {
		t.Error(r)
	}


}
