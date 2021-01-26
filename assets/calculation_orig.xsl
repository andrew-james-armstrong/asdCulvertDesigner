<?xml version="1.0" encoding="UTF-8" ?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" >
	<xsl:output method="html"/>

	<xsl:template match="/">
		<html>
			<head>
				<title>Single Cell Box Culvert Design</title>
				<link href="./style.css" rel="stylesheet" type="text/css" />
			</head>
			<body bgcolor="#ffffff">
         <img src="./asd_logo.svg" border="0" width="96" height="96" ALT="Calculation Sheet"/>
                <xsl:apply-templates />
			</body>
		</html>
    </xsl:template>

  <xsl:template match="Calculation">
    <div class="Title">
        <h1>&title1;&nbsp;-&nbsp;&title2;</h1>
        <h2>&Inp;</h2>
           <xsl:value-of select="text()"/>
    </div>
    <div>
        <a href="culvertserver.html">Change input data</a>
        <xsl:apply-templates select="Culvert"/>
    </div>
  </xsl:template>

    <xsl:template match="Culvert">
      <div class="Title">
          <xsl:value-of select="text()"/>
      </div>
      <div>
          <h2>Culvert Name</h2>
          <xsl:apply-templates select="Project"/>
      </div>
  </xsl:template>
</xsl:stylesheet>
