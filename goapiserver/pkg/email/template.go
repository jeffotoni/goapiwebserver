package email

import "strings"

func TemplateEmail(Nome, Projeto string) (tmpl string) {

	Projeto = strings.Replace(Projeto, `"`, "", -1)
	Nome = strings.Replace(Nome, `"`, "", -1)

	tmpl = `<table width="650" cellpadding="0" cellspacing="0" border="0" align="center" style="font-family: Arial,sans-serif; font-size: 14px; color: #2b2722; border: 1px solid #F6F6F6; line-height: 1.5">
	<tr>
	  	<td style="padding:18px 18px 30px;text-align:left;" colspan="3"><span style="font-weight: normal; color: #000; font-size: 35px;">api</span><span style="font-weight: bold; color: #000; font-size: 40px;">Go Server</span></td>
	</tr>
	<tr>
    	<td colspan="3" style="color:#49B5E2;padding: 0 30px;font-size: 21px;text-align:left;">
    		Ol&aacute;, ` + Nome + `<br />
    	</td>
	</tr>
	<tr>
		<td colspan="3">
			<div style="padding: 10px 30px 30px; font-size: 18px; line-height: 1.5; color: #424242;">
				Foi executado um endpoint na Api Server. <br />
				Voc&ecirc; poder&aacute; visualizar os logs de toda execu&ccedil;&atilde;o.<br />'
			</div>
		</td>
	</tr>
	<tr>
    	<td colspan="3" style="color:#a9a9a9;padding: 10px 30px 30px;font-size: 14px;text-align:center;">
    		` + Projeto + `
    	</td>
	</tr>
	</table>`

	return
}
