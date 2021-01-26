<?xml version="1.0" encoding="UTF-8" ?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:output method="html"/>

<xsl:template match="Dates">
<tr bgcolor="#d0d088">
<td width="20%" class="boldTitles">
<xsl:value-of select="." />
</td>
<xsl:apply-templates select="Week" />
</tr>
</xsl:template>

<xsl:template match="WeekDate">
<tr bgcolor="#d0d088">
<td width="30%" class="boldTitles">
<xsl:value-of select="." />
</td>
<xsl:apply-templates select="Week" />
</tr>
</xsl:template>

<xsl:template match="Resource" mode="summary">
<tr>
<td bgcolor="#ffd0ff" height="20pt" valign="top" class="SmallText">
<xsl:value-of select="text()" />
</td>
<xsl:apply-templates select="WorkSum" />
</tr>
</xsl:template>

<xsl:template match="Resource" mode="detail">
<tr>
<td height="10pt" width="20%"></td>
<xsl:apply-templates select="//Division/Summary/Dates" />
</tr>
<tr>
<td bgcolor="#ffd0ff" height="20pt" valign="top" class="SmallText">
<xsl:value-of select="text()" />
</td>
<xsl:apply-templates select="WorkSum" />
<xsl:apply-templates select="Project" />
</tr>
</xsl:template>

<xsl:template match="Project">
<tr bgcolor="#d0d088">
<td width="10%" class="boldTitles" style="font-size:6pt" height="20pt" >
<xsl:value-of select="text()" />
</td>
<xsl:apply-templates select="Work" mode="detail" />
</tr>
</xsl:template>

<xsl:template match="WorkSum">
<xsl:apply-templates select="Work" mode="summary" />
</xsl:template>

<xsl:template match="Dates/Week">
<td class="boldTitles" style="font-size:6pt" valign="top" >
<xsl:value-of select="@Date" />
</td>
</xsl:template>

<xsl:template match="WeekDate/Week">
<td class="boldTitles" valign="top" >
<xsl:value-of select="@Date" />
</td>
</xsl:template>

<xsl:template match="Resource/Week">
<xsl:apply-templates select="@Utilisation" />
</xsl:template>

<xsl:template match="Work" mode="summary">
<xsl:apply-templates select="@Util" mode="summary"/>
</xsl:template>

<xsl:template match="@Util" mode="summary" >
<xsl:choose>
<xsl:when test="number(number(.) = 0)" >
	<td bgcolor="#ffffe0" class="redText" style="font-size:6pt" >
	</td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 50)" >
	<td bgcolor="#f0f0f0" class="amberText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 100)" >
	<td bgcolor="#80b080" class="blueText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 120)" >
	<td bgcolor="#B0FFB0" class="blueText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 140)" >
	<td bgcolor="#90FF90" class="blueText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 160)" >
	<td bgcolor="#FFe0e0" class="blueText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 180)" >
	<td bgcolor="#ffa0a0" class="blueText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 290)" >
	<td bgcolor="#ff8080" class="whiteText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:when test="number(number(.) &gt; 290)" >
	<td bgcolor="#ff2020" class="whiteText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:when>
<xsl:otherwise>
	<td bgcolor="#404040" class="whiteText" style="font-size:6pt" >
	<xsl:value-of select="." />
	</td>
</xsl:otherwise>
</xsl:choose>
</xsl:template>

<xsl:template match="Work" mode="detail">
<xsl:apply-templates select="@Util" mode="detail"/>
</xsl:template>

<xsl:template match="@Util" mode="detail" >
<xsl:choose>
<xsl:when test="number(number(.) = 0)" >
    <td bgcolor="#ffffff" >
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 10)" >
    <td bgcolor="#f0f0f0" class="redText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 20)" >
    <td bgcolor="#d0d0d0" class="redText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 30)" >
    <td bgcolor="#b0b0b0" class="redText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 40)" >
    <td bgcolor="#909090" class="whiteText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 50)" >
    <td bgcolor="#80A080" class="amberText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 60)" >
    <td bgcolor="#80B080" class="blueText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 70)" >
    <td bgcolor="#70C070" class="blueText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 80)" >
    <td bgcolor="#90C090" class="blueText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 90)" >
    <td bgcolor="#A0c0A0" class="blueText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 120)" >
    <td bgcolor="#80ff80" class="blueText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 140)" >
    <td bgcolor="#ffe0e0" class="greenText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 160)" >
    <td bgcolor="#ffa0a0" class="greenText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &lt; 180)" >
    <td bgcolor="#ff6060" class="whiteText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:when test="number(number(.) &gt; 250)" >
    <td bgcolor="#ff2020" class="whiteText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:when>
<xsl:otherwise>
    <td bgcolor="#ff4040" class="whiteText" style="font-size:6pt" >
    <xsl:value-of select="." />%
    </td>
</xsl:otherwise>
</xsl:choose>
</xsl:template>

<xsl:template match="@Date">
<td>
<xsl:apply-templates select="." />
</td>
</xsl:template>

<xsl:template match="@JobNo">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@Name">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@Start">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@Finish">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@Duration">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@Bridge">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@Road">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@Rail">
<xsl:value-of select="." />
</xsl:template>

<xsl:template match="@All">
<xsl:value-of select="." />
</xsl:template>
</xsl:stylesheet>
