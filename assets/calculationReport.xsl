<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
	<xsl:output method="html"/>
	
	<xsl:template match="/" >
        <html>
			<head>
				<title>Single Cell Box Culvert Design</title>
				<link href="./assets/style.css" rel="stylesheet" type="text/css"/>
				<link rel="mask-icon" href="./assets/icon.png" sizes="60x60" color="red"></link>
				<link rel="apple-touch-icon" href="./assets/icon.png" sizes="60x60"></link>
			</head>
			<body bgcolor="#d0d0ff">
                <xsl:apply-templates select="Calculation"/>
            </body>
        </html>
    </xsl:template>
	
	<xsl:template match="Meta/Version">
		<xsl:value-of select="text()"/>
	</xsl:template>
	<xsl:template match="Meta/Date">
		<xsl:value-of select="text()"/>
	</xsl:template>
	<xsl:template match="Meta/Name">
		<xsl:value-of select="text()"/>
	</xsl:template>

	<xsl:template match="Calculation">
	        <h1>Single Cell Culvert Designer</h1>
			<p class ="copy">Designer <xsl:apply-templates select="Meta/Version"/> Date <xsl:apply-templates select="Meta/Date"/> by <xsl:apply-templates select="Meta/Name"/> All rights reserved</p>
	        <xsl:apply-templates select="Project"/>
	        <xsl:apply-templates select="Nodes"/>
	        <xsl:apply-templates select="Elements"/>
	        <xsl:apply-templates select="Results"/>
	        <xsl:apply-templates select="Charge"/>
	</xsl:template>

	<xsl:template match="Project">
	    <h1>Project - <xsl:value-of select="./project_name"/> - Design calculations</h1>
	    <p class="detail">This calculation documents the design of a single cell concrete box culvert that is buried below a road or railway. The forces on the structure are derived from the input parameters provided by the user. The design has been carried out in accordance with the Eurocode series of standards using the UK National Appendices for locale specific parameters where appropriate. In particular:<ol>BS EN 1990 - General Principles</ol><ol>BS EN 1991 Parts 1&amp;2 - Design of Concrete Bridges</ol></p>

      <table width="80%" border="0" bgcolor="a0a0ff">
          <tr><xsl:apply-templates select="project_name"/></tr>
          <tr><xsl:apply-templates select="structure_name"/></tr>
          <tr><xsl:apply-templates select="structure_reference"/></tr>
          <tr><xsl:apply-templates select="structure_number"/></tr>
		  <tr><td width="40%"><p class="detail">Type of Structure:</p></td><td><p class="detail">Single Cell Box Culvert</p></td></tr>
      </table>
	  <img src="./assets/1_cell_box_section.svg" width="800em" height="600em"/>
	</xsl:template>

    <xsl:template match="project_name">
        <td width="40%"><p class="detail">Name of overall Project:</p></td><td><p class="detail"><xsl:value-of select="text()"/></p></td>
    </xsl:template>

    <xsl:template match="structure_name">
          <td width="40%"><p class="detail">Structure Name:</p></td><td><p class="detail"><xsl:value-of select="text()"/></p></td>
    </xsl:template>

    <xsl:template match="structure_reference">
      <td width="40%"><p class="detail">Structure Reference:</p></td><td><p class="detail"><xsl:value-of select="text()"/></p></td>
    </xsl:template>

<xsl:template match="structure_number">
  <td width="40%"><p class="detail">Structure Identification Number:</p></td><td><p class="detail"><xsl:value-of select="text()"/></p></td>
</xsl:template>

<xsl:template match="Nodes">
    <h3>Model Node Co-ordinate List</h3>
	<p class="detail">The list of coordinates below is calculated following the mid depth of the structure elements. Node positions have been chosen to capture the key critical parts of the structure for design forces.</p><p class="detail"> From these coordinates, element dimensions and properties will be chosen for development of dead loads and other loading components.</p>
    <table border="0" width="50%" bgcolor="#fafafa">
        <tr><th width="20%">Node Number</th><th width="25%">x</th><th width="25%">y</th><th width="25%">z</th></tr>
        <xsl:apply-templates select="point"/>
    </table>
</xsl:template>

<xsl:template match="Elements">
    <h3>Structure Element List</h3>
    <table border="0" width="100%" bgcolor="#ddddff">
        <tr><th width="5%">Element Number</th><th width="5%">Start Node</th><th width="5%">Finish Node</th><th width="10%">Breadth</th><th width="10%">Height</th><th width="5%">Concrete Grade</th><th width="5%">Reinforcement Grade</th><th width="10%">Shear Link Allowance</th><th width="10%">Minimum Cover</th><th width="10%">Cover Tolerance</th><th width="20%">Reinforcement Layers</th></tr>
    	<xsl:apply-templates select="Element"/>
    </table>
</xsl:template>

<xsl:template match="point">
    <tr>
	<td class="detail">
        <b><xsl:value-of select="@n"/></b>
    </td>
    <td class="detail">
        <xsl:value-of select="x"/>
    </td>
    <td class="detail">
        <xsl:value-of select="y"/>
    </td>
    <td class="detail">
        <xsl:value-of select="z"/>
    </td>
	</tr>
</xsl:template>

<xsl:template match="@n">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="Element">
	<tr>
		<td class="detail"><b><xsl:value-of select="@n"/></b></td>
		<td class="detail"><xsl:value-of select="start"/></td>
		<td class="detail"><xsl:value-of select="end"/></td>
		<xsl:apply-templates select="Rectangular_concrete_section"/>
	</tr>
</xsl:template>

<xsl:template match="start">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="end">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="Rectangular_concrete_section">
	<td class="detail"><xsl:apply-templates select="Breadth"/></td>
	<td class="detail"><xsl:apply-templates select="Height"/></td>
	<td class="detail"><xsl:apply-templates select="ConcreteGrade"/></td>
	<td class="detail"><xsl:apply-templates select="ReinforcementGrade"/></td>
	<td class="detail"><xsl:apply-templates select="ShearLinkAllowance"/></td>
	<td class="detail"><xsl:apply-templates select="MinimumCover"/></td>
	<td class="detail"><xsl:apply-templates select="CoverTolerance"/></td>
	<td class="detail">
		<xsl:apply-templates select="ReinforcementLayers"/>
	</td>
</xsl:template>

<xsl:template match="Breadth">
    	<xsl:value-of select="." />
</xsl:template>
<xsl:template match="Height">
    	<xsl:value-of select="." />
</xsl:template>
<xsl:template match="ConcreteGrade">
    	<xsl:value-of select="." />
</xsl:template>
<xsl:template match="ReinforcementGrade">
    	<xsl:value-of select="." />
</xsl:template>
<xsl:template match="ShearLinkAllowance">
    	<xsl:value-of select="." />
</xsl:template>
<xsl:template match="MinimumCover">
    	<xsl:value-of select="." />
</xsl:template>
<xsl:template match="CoverTolerance">
    	<xsl:value-of select="." />
</xsl:template>

<xsl:template match="ReinforcementLayers">
	<table border="0" bgcolor="#eeeeff">
		<tr><th class="detail">Layer</th><th class="detail">Primary Bar</th><th class="detail">Secondary Bar</th><th class="detail">Spacing</th><th class="detail">Depth below NA</th></tr>
    	<xsl:apply-templates select="Layer" />
	</table>
</xsl:template>

<xsl:template match="Layer">
    <tr>
	<td class="detail"><xsl:apply-templates select="Name"/></td>
	<td class="detail"><xsl:apply-templates select="LargeBar" /></td>
    <td class="detail"><xsl:apply-templates select="SmallBar" /></td>
    <td class="detail"><xsl:apply-templates select="Spacing" /></td>
    <td class="detail"><xsl:apply-templates select="Level" /></td>
	</tr>
</xsl:template>

<xsl:template match="Name">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="LargeBar">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="SmallBar">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="Spacing">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="Level">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="Results">
    <h3>Calculated Analysis Results</h3>
    <table border="0" width="100%" bgcolour="#fafaff">
            <xsl:apply-templates select="Parameter"/>
    </table>
</xsl:template>

<xsl:template match="Parameter">
    <tr>
        <td class="detail" width="10%"><xsl:apply-templates select="@ref"/></td>
        <td class="detail" width="40%"><xsl:apply-templates select="text()"/></td>
        <td class="detail"><xsl:apply-templates select="@value"/></td>
        <td class="detail"><xsl:apply-templates select="@unit"/></td>
    </tr>
</xsl:template>

<xsl:template match="@ref">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="@value">
    <xsl:value-of select="." />
</xsl:template>

<xsl:template match="@unit">
    <xsl:if test=". = 'm'">m</xsl:if>
    <xsl:if test=". = 'm^2'">m<sup>2</sup></xsl:if>
    <xsl:if test=". = 'mm'">mm</xsl:if>
    <xsl:if test=". = 'm^3'">m<sup>3</sup></xsl:if>
    <xsl:if test=". = 'm^4'">m<sup>4</sup></xsl:if>
    <xsl:if test=". = 'N/sq mm'">N/mm<sup>2</sup></xsl:if>
    <xsl:if test=". = 'kNm'">kNm</xsl:if>
    <xsl:if test=". = 'kN/sq m'">kN/m<sup>2</sup></xsl:if>
</xsl:template>

<xsl:template match="Charge">
    <h3>Design and Analysis Charges</h3>
    <p>The cost of this analysis is based on the time required to input, analyse, design and present the results.</p>
    <p>The client that requested this analysis is <xsl:value-of select="@Client"/></p>
    <p>The Account Number that has been charged is <xsl:value-of select="@Account"/></p>
    <p>The amount charged to the client account for this analysis is £<xsl:value-of select="@ItemCost"/>.</p>
</xsl:template> 
</xsl:stylesheet>
