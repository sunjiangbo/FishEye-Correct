GOOF----LE-8-2.03M      ] Ĩ 4 h      ] g  guileĪ	 Ī	g  process-use-modulesĪ	 Ī	 Ī	g  	aisleriotĪ	g  	interfaceĪ	 Ī		 Ī	
g  apiĪ	
 Ī	 Ī	g  ice-9Ī	g  formatĪ	 Ī	 Ī	g  reserve-sizeĪ	g  tableau-sizeĪ	g  build-foundation-in-suitĪ	g  select-baseĪ	g  
max-redealĪ	g  auto-fill-tableauĪ	g  fill-from-reserveĪ				
 Ī				 Ī				
 Ī				
 Ī			
 Ī		
	
 Ī		
		
 Ī	 Ī	 g  
variationsĪ	!g  _Ī	"f  TerraceĪ	#f  General's PatienceĪ	$f  Falling StarsĪ	%f  SignoraĪ	&f  RedheadsĪ	'f  Blondes and BrunettesĪ	(f  WoodĪ	)g  variation-namesĪ	*g  current-variationĪ	+g  BASE-VALĪ	,g  call-with-deferred-observersĪ	-, Ī	., Ī	/g  module-export!Ī	0/ Ī	1/ Ī	2g  current-moduleĪ	32 Ī	42 Ī	5+ Ī	6g  variable-listĪ	7g  stock-lockedĪ	87 Ī	9g  stockĪ	:g  wasteĪ	;g  reserveĪ	<									
 Ī	=g  
foundationĪ	>g  tableauĪ	?g  initialize-playing-areaĪ	@g  set-ace-lowĪ	Ag  make-standard-double-deckĪ	Bg  shuffle-deckĪ	Cg  add-normal-slotĪ	Dg  reverseĪ	Eg  DECKĪ	Fg  add-blank-slotĪ	Gg  add-extended-slotĪ	Hg  rightĪ	Ig  add-carriage-return-slotĪ	Jg  build-tableau-slotsĪ	Kg  deal-reserve-cardsĪ	Lg  deal-tableau-cardsĪ	Mg  deal-cards-face-upĪ	Ng  	get-valueĪ	Og  get-top-cardĪ	Pg  do-auto-dealĪ	Qg  give-status-messageĪ	Rg  calculate-scoreĪ	Se  4.1Ī	Tg  new-gameĪ	Ug  SLOTSĪ	Vg  downĪ	Wg  HORIZPOSĪ	Xg  set-statusbar-messageĪ	Yg  string-appendĪ	Zg  get-stock-no-stringĪ	[g  get-redeals-stringĪ	\f     Ī	]g  get-base-stringĪ	^f  Base Card: AceĪ	_f  Base Card: JackĪ	`f  Base Card: QueenĪ	af  Base Card: KingĪ	bf   Ī	cf  Base Card: ~aĪ	dg  number->stringĪ	ef  Redeals left:Ī	ff   Ī	gg  FLIP-COUNTERĪ	hf  Stock left:Ī	ig  lengthĪ	jg  	get-cardsĪ	kg  aceĪ	lg  kingĪ	mg  descending-values?Ī	ng  calculate-score-helperĪ	og  
set-score!Ī	pg  empty-slot?Ī	qg  do-auto-fill-tableauĪ	rg  	dealable?Ī	sg  do-deal-next-cardsĪ	tg  memberĪ	ug  button-pressedĪ	vg  move-n-cards!Ī	wg  or-mapĪ	xg  complete-transactionĪ	yg  
droppable?Ī	zg  button-releasedĪ	{g  	get-colorĪ	|g  get-suitĪ	}g  
flippable?Ī	~g  
flip-stockĪ	g  button-clickedĪ g  remove-cardĪ g  auto-play-to-foundationĪ g  button-double-clickedĪ f  	somethingĪ f  the foundationĪ g  hint-start-foundationĪ g  	hint-moveĪ g  hint-slot-to-foundationĪ g  hint-slots-to-foundationĪ g  get-rankĪ g  droppable-on-foundationĪ g  buildable-on-tableau-helperĪ g  buildable-on-tableauĪ g  hint-tableau-build-helperĪ g  hint-tableau-buildĪ g  hint-waste-to-tableauĪ f  Deal a new card from the deckĪ g  	hint-dealĪ g  get-hintĪ g  	get-scoreĪ g  game-wonĪ g  game-continuableĪ g  get-variation-optionsĪ g  appendĪ g  begin-exclusiveĪ  Ī g  end-exclusiveĪ  Ī g  get-optionsĪ g  list-refĪ g  set-variationĪ g  apply-optionsĪ  g  timeoutĪ Ąg  set-featuresĪ Ēg  droppable-featureĪ Ģg  dealable-featureĪ Īg  
set-lambdaĪC 5     hĀ<  	  ] 4	 >  "  G  	R		RRR
RRR R4!i"54!i#54!i$54!i%54!i&54!i'54!i(5 )R
*R
+R4.145        h   =   ] 45 6   5       g  filenamef  terrace.scm
	6
 		
   C>  "  G  +6i6i7R4.148  h   =   ] 45 6   5       g  filenamef  terrace.scm
	;
 		
   C>  "  G  76i6i
9R:R	;R<=R>R?@ABCDE9:FGH;I=>JKL+M9=NOPQRS !h  ņ  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  445>  "  G  4	>  "  G  4
>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G   4>  "  G  45 4>  "  G  $  4	>  "  G  
 "  ;4 >  "  G  4455 4>  "  G  4>   "  G  4>   "  G  4>   "  G  	  C é      g  filenamef  terrace.scm
	C
		D			E		#	G		3	H		C	J		F	J		N	J	"	S	J		\	K		_	K		a	K		f	K		o	L			M	 	M	 	M	 	M	 	N	 Ī	P	 §	P	 Đ	P	 Ū	P	 ·	Q	 š	Q	 ž	Q	 Á	Q	 Ę	R	 Í	R	 Ï	R	 Ô	R	 Ý	S	 ā	S	 â	S	 į	S	 ð	T	 ó	T	 õ	T	 ú	T		U		U		U		U		V		V		V	 	V	)	W	,	W	.	W	3	W	<	X	L	Z	N	Z	O	[	a	\	i	\	j	^		`		a		b		c	Ī	c	-§	c	'Ž	c	ĩ	d	ļ	d	'―	d	5ŋ	d	'Á	d	Ã	d	Ä	e	Ø	g	č	h	ø	i		k	
	k	 N	
  g  nameg  new-game CTRU>GV>WJ 	       hH   ų   ] 
$  C $  *45$  	 $   6CCC    ņ       g  count
		D  g  filenamef  terrace.scm
	n
		o		
	o			p			p			o			q			q		 	q	#	"	q		&	o		.	r	'	/	r	"	0	r		2	r		7	o		<	s		>	s	 		D  g  nameg  build-tableau-slots CJRM9;K        h(   °   ] 
$  C4 5$   6C   Ļ       g  count
		%  g  filenamef  terrace.scm
	u
		v		
	v			w			w	!		w			v		!	x		#	x	 
		%  g  nameg  deal-reserve-cards CKRM9L   h8   Ó   ]
$  C (  C4  5$  
 6C       Ë       g  slots
		1 g  count		1  g  filenamef  terrace.scm
	z
		{		
	{			}			}	'		}	!	!	}		%	{		*	~		-	~	'	/	~	 		1	  g  nameg  deal-tableau-cards CLRXYZ[\]  h       ] 445 45 45 56           g  filenamef  terrace.scm
 
	 		 	(	 	(	 	(	 	(	 		 	 			
  g  nameg  give-status-message CQR+!^_`abcd        hp   M  ] "  >$  6	$  6	$  6	$  6C$   	$  4	54
56"ĸĸ"ĸĸE      g  filenamef  terrace.scm
 
	
 			 		 		 			 			 		! 		# 			( 			, 		0 		2 			7 			; 		? 		A 			C 		D 		H 		L 		Q 		U 			Y 		] 		_ 		` 	'	h 		 		p
  g  nameg  get-base-string C]R+bY\!efdg        h@   ë   ]  $  "  
 $  C454	
56      ã       g  t
		  g  filenamef  terrace.scm
 
	 	
	 		 		 		  		% 		& 		* 		, 		. 	.	/ 		6 	%	8 		: 	 		:
  g  nameg  get-redeals-string C[R+bY!hfdij9  h0   Ó   ] 
$  C45444	
5556     Ë       g  filenamef  terrace.scm
 
	 		
 		 		 		 		 		 	&	 		 	%	 	-	' 	%	) 		+ 	 		+
  g  nameg  get-stock-no-string CZRkl  h(   Á   ]
 $  C $  CC    đ       g  a
		$ g  b		$ g  t			$  g  filenamef  terrace.scm
 Ą
	 Ē		 Ē		 Ē		 Ģ		 Ģ		! Ī	 		$	  g  nameg  descending-values? CmRnij  h(   Ũ   ] (  C 44 556      Ï       g  slots
		" g  acc		"  g  filenamef  terrace.scm
 Ķ
	 §		 Đ		 Đ	1	 Đ	9	 Đ	D	 Đ	9	 Đ	1	  Đ	*	" Đ	 		"	  g  nameg  calculate-score-helper CnRon=    h   l   ] 4
56  d       g  filenamef  terrace.scm
 Ŧ
	 Ž		 Ž	 		
  g  nameg  calculate-score CRRp:9M;q       h°   e  ] (  C4 5$  "  R45$  +45$  "  q4  >  "  \G  "  U4  >  "  @G  "  9$  /45$  "ĸĸ"  4  >  "  G  "  "ĸĸu"    6]      g  slots
	 °  g  filenamef  terrace.scm
 Ū
	 Ŋ		 ą		 ą		 ą		 ą		 ī		% ē		& ķ		0 ē		5 ·		< ·	9	? ·	3	D ·		Q ĩ		X ĩ	9	[ ĩ	3	` ĩ		l ē		s ē	5	} ē	  ģ	  ģ	;  ģ	5  ģ	 Ū ļ	# ° ļ	 	 °  g  nameg  do-auto-fill-tableau CqR+q>p:rs 	   h   "  ]
  $  "  $$  "  4>  "  G   
  $  "  <45$  "  %45 $  "  4>   "  G   C            g  t
		9 g  t		6 g  t
	=  g  t	Q  g  t		c   g  filenamef  terrace.scm
 š
	 ŧ		 ŧ		 ž		 ŧ		" ―		= ū		= ū		J ŋ		Q ŋ		Q ū		^ Ā		c Ā		c ū		p Á	 	 
  g  nameg  do-auto-deal CPR+t>i;:   hH   é   ]
$   64 5$  
45C $  C $  CC       á       g  slot-id
		A g  	card-list		A  g  filenamef  terrace.scm
 Ä
	 Å			
 Å		 Æ			 Į			 Å		! Č		( Č			. É			2 Å		9 Ë			= Å	 		A	  g  nameg  button-pressed CuRv+NOq>:wp7P  h   W  ]4 >  "  G  
$  "   4455 4>  "  G   $  "  
4	5$   
"   4>   "  G  CO      g  
start-slot
	  g  	card-list	  g  end-slot		  g  t			L g  t		Q	k  g  filenamef  terrace.scm
 Ï
	 Ð		 Ņ		 Ņ		 Ņ		* Ō		- Ō	'	5 Ō		7 Ō		8 Ó		Q Ô		Q Ô		_ Õ		h Õ		o Ô		r Ö		w Ũ	 	 	  g  nameg  complete-transaction CxRyx      h    ŧ   ]4 5$  
 6C   ģ       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  terrace.scm
 Ú
	 Û		 Û		 Ü	 			  g  nameg  button-released CzR+t>=9i;p:{OmN|        hH  §  ]
$  4 5$  6C $  C45$  45$  r $  C $  C45$   	C4
54
455$  C44455455$  4455CCC45$  ~45$  45C$  454455"  4
54
455$  ,44544555$  45CCCC       g  
start-slot
	G g  	card-list	G g  end-slot		G  g  filenamef  terrace.scm
 Þ
	 ß			
 ß		 ā		 ā			 á		& â			* ß		- ä			9 ß		; å		B å		F å			K æ		O å			V į		Z å			] č		g č		l é		n ę		s ę	*	u ę		v ë		y ë	*  ë	  ę	  ę	  ė	  ė	+  ė	6  ė	+  í	+  í	6  í	+ Ą ė	 Ĩ ę	 Ķ î	 Đ î	* ą î	 ī î	 ĩ î	 ŧ ï		 Į ß	 Č ð	 Ō ð		 Õ ņ	 Ú ņ	$ Ü ņ	 Ý ņ	 ä ō	 å ó	 ę ó	# ė ó	 í ô	 ð ô	# ø ô	 ų ó	 þ õ	 õ	) õ	 ö		 ö	) ö	 õ	 õ	 ō	 ũ	 ũ	&  ũ	1" ũ	&# ø	&& ø	1. ø	&0 ũ	4 ō	5 ų	: ų	%< ų	? ų	@ ų	 R	G	  g  nameg  
droppable? CyR+}9:gp7 
       hP   Ė   ]
$  C45$  2  $  "  45 $    $   C	CCC Ä       g  t
		6 g  t
	<	K  g  filenamef  terrace.scm
 ü
	 ý		
 ý		 þ		 ý		 		 		-		: ý		<		J	 		O
  g  nameg  	dealable? CrRr~9:wp>7 
  hH   ū   ]45 $  845$  &45  $  "   	 $  CCCC   ķ       g  t
	%	9  g  filenamef  terrace.scm

											%		%		5			=	 		E
  g  nameg  do-deal-next-cards CsR9s       h       ] $  45 $  CCC             g  
start-slot
		  g  filenamef  terrace.scm

								 		  g  nameg  button-clicked CRpyOx       hP     ](  C4 5$  C4 4 5 5$   4 5 6 6           g  
start-slot
		K g  	end-slots		K  g  filenamef  terrace.scm

										(	&	"	)	C	+		/		4	2	=	,	@	L	B		I	/	K	 		K	  g  nameg  auto-play-to-foundation CR=    h      ] 6      z       g  
start-slot
		
  g  filenamef  terrace.scm

	
	 		
  g  nameg  button-double-clicked CR+!       h    ą   ] 
$  	4545 CCĐ       g  filenamef  terrace.scm

			
											 				 		 
  g  nameg  hint-start-foundation CRpyO      hH     ](  C4 5$  C4 4 5 5$  
 6 6      ĸ       g  
start-slot
		B g  	end-slots		B  g  filenamef  terrace.scm

	 		!		 		"		"	(	&"	"	)"	C	+"		/"		7#	#	9#		@$	/	B$	 		B	  g  nameg  hint-slot-to-foundation CR=    h(   Ę   ]	 (  C4 5$  C 6  Â       g  start-slots
		& g  t		&  g  filenamef  terrace.scm
&
	'		(		(	$	(		(		$)	%	&)	 			&  g  nameg  hint-slots-to-foundation CR+     h      ] $  	 C C    z       g  value
		  g  filenamef  terrace.scm
+
	,		,		-	 		  g  nameg  get-rank CRy   h0   ó   ](  C4 5$  C 6    ë       g  
start-slot
		, g  cards		, g  	end-slots			, g  t			,  g  filenamef  terrace.scm
0
	1		2		2	(	2		2		*3	5	,3	 			,	  g  nameg  droppable-on-foundation CR=>N 	     h   Å  ](  "  $  (  "  C4  5$  "  $  C4 5$  1 $  "  44554 56C   ―      g  
start-slot
	  g  	num-cards	  g  cards		  g  acc		  g  t		;  g  target-slot		R   g  filenamef  terrace.scm
7
	8		8		8		"<		)<	5	,<	/	0<		4<		;8		G>		N>	A	R>		R>		Z?		aB		dC		jD		qD	#	tD	-	yD	8	{D	-	}D	#	~D	F D	 @	 	 	  g  nameg  buildable-on-tableau CRpy      hP   5  ](  C45$  "  4  5$  "  $  C 6    -      g  
start-slot
		L g  card		L g  	end-slots			L g  t		6	L  g  filenamef  terrace.scm
E
	F		G		G	"	G		G		H		%H	'	(H	3	*H		.G		1I		6G		JJ	8	LJ	 		L	  g  nameg  buildable-on-tableau-helper CR        h@     ]
 (  C4 5 $  $  "  "  6          g  start-slots
		< g  acc		< g  hint			<  g  filenamef  terrace.scm
L
	M		N		P		P	-	P		P		R		!S		$S	 	'S	+	(S		,S		<Q	 		<	  g  nameg  hint-tableau-build CRpj    h    Ą   ]4 5$  C 4 56         g  
start-slot
		  g  filenamef  terrace.scm
X
	Y		Y		Z	*	Z	 		  g  nameg  hint-tableau-build-helper CRp:yO    h@   í   ] (  C45$  C445  5$  
 6 6å       g  	end-slots
		@  g  filenamef  terrace.scm
\
	]		^		]		_		_	#	&_		)_	9	+_		/_		7`		9`		>a	"	@a	 		@  g  nameg  hint-waste-to-tableau CRr!      h       ] 45 $  
45 CC              g  filenamef  terrace.scm
c
	d		d		e		e		e		e	 		
  g  nameg  	hint-deal CR;=>:+ hx   0  ]45   $   C45  $   C45  $   C45  $   C4	5  $   C4
5  $   C6 (      g  t
		x g  t
		x g  t
	-	x g  t
	A	x g  t
	W	x g  t
	i	x  g  filenamef  terrace.scm
g
	h		h		i		h		'j		-h		9k		Ah		Ml		Ul	"	Wl		Wh		cm		ih		xn	 		x
  g  nameg  get-hint CR       h   e   ] 45 	hC     ]       g  filenamef  terrace.scm
p
	q		
q	 		
  g  nameg  game-won CRQR    h8      ] 4>   "  G  4>   "  G  45 $  C6               g  filenamef  terrace.scm
s
	t		u		#v		+v		1w	 		1
  g  nameg  game-continuable CR*     h(   á   ] (  C  4 5C      Ų       g  names
		" g  index		"  g  filenamef  terrace.scm
y
	z			{		|		|		|		}		}	#	}	/	 }		!|	 		"	  g  nameg  get-variation-options CR)      h   z   ] 4
56      r       g  filenamef  terrace.scm

		
		
		
		 		
  g  nameg  get-options CR*   hÐ   G  ] (  C$  Ą $  4 
5 $  ~4 5 $  i4 	5 $  S4 	5 $  =4 	5 $  '4 	5 $  4 	5 	"  "  "  "  "  "  "  "  $  C
 6   ?      g  vars
	 Í g  options	 Í g  index		 Í g  t	 ĩ Í  g  filenamef  terrace.scm

																"	(	%		'		,		-		2	(	5		7		<		=	*	B	4	F	*	H		M		N		S	'	W		Y		^		_		d	&	h		j		o		p	#	u	-	y	#	{	 	 	# 	- 	# 	 ĩ	 Å	 Č	  Ë	. Í	 /	 Í	  g  nameg  set-variation CR       h      ] 
6    x       g  options
		  g  filenamef  terrace.scm

					 		  g  nameg  apply-options CR   h   R   ] C    J       g  filenamef  terrace.scm

 		
  g  nameg  timeout C R4ĄiĒiĢi>  "  G  ĪiTiuiziiiiiiii iyiri6             g  filenamef  terrace.scm		
	 	
	%	
	)	
	-	
	1	
	5	
	9	
	;			>	
	?	&		E	&	
	G	&		H	(		N	(	
	P	(		Q	*		W	*	
	Y	*		Z	,		`	,	
	b	,		c	.		i	.	
	k	.		l	0		r	0	
	t	0		u	2		{	2	
	}	2	 	$	 	#
 	4
 	6
	;
	=
	>
	?
	@		@
 	A	Ģ	A
ô	C
\	n
	S	u

v	z
P 
8 
 
ķ 
ī Ą
Ę Ķ
_ Ŧ
 Ū
m š
ž Ä
Â Ï
ķ Ú
Ú Þ
! ü
"I
#
$Ī
%M
&<
'Š
(ĩ&
)e+
*0
-7
.ļE
0"L
0üX
2H\
3c
4Ûg
5hp
6Cs
7dy
8
;R
;û
<e
<f
<đ
 W	<đ
   C6 