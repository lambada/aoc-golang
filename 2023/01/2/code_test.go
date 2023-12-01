package y2023d01p02

import (
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example 1",
			input: "two1nine",
			want:  29,
		},
		{
			name:  "Example 2",
			input: "eightwothree",
			want:  83,
		},
		{
			name:  "Example 3",
			input: "abcone2threexyz",
			want:  13,
		},
		{
			name:  "Example 4",
			input: "xtwone3four",
			want:  24,
		},
		{
			name:  "Example 5",
			input: "4nineeightseven2",
			want:  42,
		},
		{
			name:  "Example 6",
			input: "zoneight234",
			want:  14,
		},
		{
			name:  "Example 7",
			input: "7pqrstsixteen",
			want:  76,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateLine(tt.input)
			if result != tt.want {
				t.Fatalf(`received %d, expected %d`, result, tt.want)
			}
		})
	}

}

func TestTotalExample(t *testing.T) {
	input := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	want := 281
	result := calculate(input)
	if result != want {
		t.Fatalf(`received %d, expected %d`, result, want)
	}
}

func TestActualPuzzle(t *testing.T) {
	rawInput := "sixsrvldfour4seven\n53hvhgchljnlxqjsgrhxgf1zfoureightmlhvvv\nfives2dznl\ntwocrqvjsix5threethree\ngtjtwonefourone6fouronefccmnpbpeightnine\nseventdtrcseveneightsevencgjgjxfpmfsix8twones\nfourthreeseven1grvhrjxklh3ninetwothree\nfourninethrnnth8\ntwo2hnxcfivejrdjxtb\nbssbrgcx86vsmqstrxsjbjeightqzhbzxqg5\nrsevenbqjfxh23rzdgcmdleightsixknpfs\njprxbcghdgxhk9x\n8czzpmvgmlchnkf\nvdrtteight85hlkninehgnblqnsdnineeight\nfivetwohlxdql43kfzzbhtncg\nlmhtwoneghgtl2\ndrtfsfhtbgfourpfpcznzrsix33\npptqkrdbeighteightsixdlsixkrlhr2\n57threetwo\nqoneight367gqbpbtbffivetlhrdjgnml2d\nninevdnzxsxone3\n3threeqzrmfndfiveeight4\nrmthreethreefourtwotkpncxpmctkxjhzxg5\nnsfrpl6lkqrrmlxhrxmngzjtzbpsjgeight\n9twobfknine57sixfour7\n7one83sixthreecllxjnphb\n82nhxcjnck1\n5zrffive\nbgdzqqpninefourjhtcjfjmxrmqhgz7fxnjrjfivetlpbddcsfs\nxmbgqfhxx1sevenkfnqztfourzbngv\n6onethree\n1przkgcft77xblxkxm\nhnjslmc96\nfourktlnine59\n7twoxthree\n9twofiveonetwo5bnl\n6cbzmrg1three7\nfbndztsthreefourthreefgj1eightrffdbpsn\n4zpqnskhnqjbjlhbbsxsmk\nseventhree1fourbvnbbsbkl5\nfive42jkbzqcnblr\n357\n543cshpxrfnnhonetkbhxtmlgczdndqjscb\n2mpftseven\n44five8nineeight\nrmgrljrljb8hxxfmdpbbvmblltxfive6mdsmm7mmpk\nnprsix6\nfive4vkqrsixmjkjps\nkfivebrlzvc1twoseveneight\nlqnjsixtwo7fourbmbxnr9dqmpbbfkfive\n1ccsvvrrvjseight9nzgnldhzndpvdzcncl\nzdzlsix9rjvninehsix4\nfsevenfour8four6three\neightrmvsq65mpvffkkhxc6seveneight\nsixrmgchvfttpbsncc5hnine\n147\n57gsdhqjbrfsskhmtcbszmjmhrxg7\nsevenfive31twomhkhttdprdnzlng7three\n3spgnkpeighttwotwo4\npsxlxdjrtmgpveightxjkxhmbeightsix2\nfive53vgqcfmgqlm5dkchcmfkclpfxntbf\n76sixmgqgftzfive4nfsqhkqvptnnshqlhd\nsevensixsevenone7xbsxvdpdtc\nnine3sixfourtwo6mpmpjb\neight4sevensixfour7dtgrvpq\n1sixnceight1bg\n79xfftvpvfour\nnoneight3\ntwoxdhkkgxhvz5three9seven8hszqdvnhr\nkbpc9kqvhonelljprfs4pzcnpfrdssfivetwo\nfour4eightseven\n5three95cqs65nine3\n72sevenrkcqmqbscp9kzjxrg8d6\n5rjbzpghninexfzvjpqmt9eight\njfjrvzcjkrkd1eight\nsix4ccvbclxkkntzmgthreercsjbznfive\nxkqdtghlshthreeninefive53three\none3fgcsrxk1\n1foursevenmvmzpvkqseven351\n858qqbkgjxqfive\nfkpd4three8sixlkttgnfzc\nbtpmrzqmsevennine3f6\n61nine\nfourone7fsfxfrlqzvbqjpccfour\nfive85sseven7jxlpl\nkzgszphcpfoursevenlzlkqmjndz4ninelnine\n4eightld\nfive8dqzlzfnine\n6zjsqxqnineplbqhxltmffour\n1fivebfoursrdcfive\n8678fourgzdhrtrtwothree\n847\nfiveeightdseven5fivefivecbdrrbnng\n5qddhnzxlhqeight2sevenln8nine9\n13kxvqlvt7xvcksdthreekfqc\n1threehqshd9two19\ntwogcphmsspeight7\n1pnktcbbpkmeight42hcgfoursix\nnineninepdxqg4xspkqjneightchqxptninegcckrbmsvk\nfnine5ninex1eightgdhpkngmgq\njlgvcn73three\ntwoljsp5\nxlf52twovcnnjglhmm\nlmonesix1kfhbpdntpthjkl\n89six6threegmfqlone\nninefbdvbhrgfivesevenzsjlt5tqkzfxb9\npvfoursixgdtvxxvlnx8kgsj\n9fpjsnheightfivermpxrkddczvzqbdbs\n92zhksstbzr\nfvxeightwoddrhpfsevenmlkp178\n2ntqvj42eight6\n46xthree26snkhxmmzknhrvsfrprch\nlvcfiveseven3plzlsvmm\nthree69\nscrkglnine1\nfivemlvcpxfcggllpnpsdjmnhhkfkbtmcfone2\n7tseveneightnmjqbsfgh\nonegbdpmtwo39nvbdl2\n532rg\n3cpjksevenonejjndff9six\n4nbtbrmtxmpcnvkjsgseven\n2brfkknktcz\nghtshqhnmvmgdps6rclhhsix\neightj4hnnsfssqqmfourgdrmgcljdnslrhmkc\nseveneight5pdzgpjvzlsxkjpflkjbzqhrfdlbskfbnbkcpl\nsix5gqjlsevenfnffpl2hrdrvrdthreeb\n137nineone81one\ngnrxrnqjdtxvcbggcmcfshz8rrkfd\nnrlccp4szxqgljj31\n622kvttthree3jfd\n8zqgbdc\nrmczlxrvjpp79sevenonefourmmjvgthree\n3fournine185sevenrgtrrpgjf\nfive282six\nthree7xcqlxd16sixktfmxslk\n6vrzsmsv8\n4twodqjxsbklhjtfthree\n3dmxljghfourzkvxq3\n4lveighteight51tfgbbfour\ncsfkthreesixxcfpdone3tvscpp8three\nvhmkkks6glrlpv\neighthnine1rsmbpndnqqgds\n7tnbhktwofive84seven\n3kjpxgrlmstnxxdnqgpgjhcttlmnineninesix5\n2eight4264\ntwoqmvnmxstnqfour8d621kxktzk\npdksthree5gzpvlmjmfourtwocfbsthree\nrtvtffd6fivefive\nonemxrpmlxsbd9drnsfive3\nzfjqcthreesixoned1sixsx\nflsevenvqctldf6\n1seventwoeight6\nxxktlvdcdtstsqjmd47onell5\n6tbvxcgddrfjnhgntrthreegx4\nrfxjvdllvzv2dshhdtjvbxvsqmq4ninegp\nfive7bjzhkmmfnhcmnfljlhb6bptflhrrvgfour\nltzlonenine5six\nnineeight49\nfqcdtgtb9sevenrnsgx4\neightlfdmv5threehrrkvhzhhxmzprz\n2dglljrllsix\nsevensixxnlb473twotwopgjxgp\n5sevenrkhk\n9eight3\n7one5four\nlh5jzqxeight84\n385nslvbjqvkxvvcprqghxccv87fpmqfhnlv\nngsjdpvv72sixmbtr\nxvslr642\nmrjhmtnhx2\ngsrbhmh5fiveoneqhhmsmkrbc3sixtwoeight\nmnxvtn3vtqhxsrfivefive57qrmhshkk\n9vkghgztxdzcvl7dlthreefmklvsmhxnqrtzzsvvq\ndjclthdptrk8fournineqvnrrlbjjhtwoeight\n79pnt9one9\nmlcxcrjzpvlkkdpmmsevenone6npd\nfivefiveninezhppshtks1kqtkss\n7nine46cxzcjhbnbg5\nfourqhbr5cqbdmvps\npqvzfrt7ftplx\n265pninenine5twofive\nfour41\n1fdgcmbc1onepcktj3seven\n5twob\ntwo1sixc7znfhkhzgzzjzf3xdbqtdspv\nvqkg7sixrntjjzc\n7one9zdnnqtbbq4gkkkdskbtlqspsmx1pphnh\nxgsix2hltpnqxktwo\ngmhmskd37nqk7q8\nonezxllhcpjfkeightsjvvpmcztdndf8seven\nninenmrljg6fivegmtzhflztxn\nnmzqxlx387eightczpsgxskfzklk9\n1fournine7cgtmr2lsvljtp2\n98lzz1sixfour1three6\none8ninetwo\n4xcrhxlvsr7\n8two2four\nninefninefour5qlxbmk51\nhhlflbbmbd29jrjrbx5879\nsix3djmgs\ncfivevgzckfqcg6ninesixfvlkpm\n449honexjqtrxveight3nzzx\nthreetwo22\nd6seventhreeeight949\nnllbxt3qspvt6sixgglvfcfivetwo\njhfrfgjjtsseveneight76\ngjninestxljz88six2\n5rsqhsx9rnd\n7three72sixjkxlnpt\ndhkfvps9\nonefive7qdvbqjxzk\n7sevenldcfd7\ntwoeightpjpkcjvvbfseven6\nsix1eight9five29\nhgdztcbpx4sixseven1fklt8six\n6six7kjlpmld1\ntwolpxfhllfllkcksvsllbtvjsll2bjq\n2nine7nine79fourfourbtqd\ntjfsmthreeoneninendd6two\n2five1\n1seven1jrgtkcfsjfbfrmsevenfourlvbgbknn\n79twolthjsjnbrthreerdmbk3\nsevennine6432\n1hjtcnjztkvjspr7\nseven84dlbb1\ntwo78threevdhgqfzt\n7bmln29onezlxtzthj7fzpvnxllz\nnldvscqn9hfnthreesevenonetwoflbrdnhq\nsixpkx7eight1ctbhdf34\nrprprpmfour3fivevxqrqrqp4tghglvssfive\ngpxjnsrsonezn1bht\n2three4ninebnfzzxzl8seven\ntwo23one2\n9sevennmxmxvzcpt3hgsprhmvtxbqvbn2three\n69threeonervrxqdfive\nscjjqgrbnnsztvmbtninethree74two2\none7nkbn6nine\nnlmdfshone2eight11sqhsftrg\nfxt4twofive3sevensevenxq\nthree6zfzpdmrpbr9289vszpcqkq\n8threebqzbpz3\nplpnsfour6\n1twonetdm\noneseventwo4x\nfourhmprhkfgdkgqznineeighteight2\n984mrrqjhzc3dsqpjvrpctvbtk\nonenjdmrjfdjznrffive73fivechpxdpnqbnpq\n4mhjzvbgnxpq1jfllblcjkbxzrt\nthreeblztlfgltlznxv9cqsjfnmgftnmscjmxpmfkleight\njtvjjrjnf2two\n8oneseven3djtdzseven3\nbctvlxxxgtp4\nonenrncdseven2foursixone\nsevenseven1\ncfhszdltcqkllpqbfive6sevenqjqphqkhhjsixbvblzhgfp\ndpzbcnbdzzninev6eightlk\neightjznzkdg53fourninebsfdzqpsix\nfour4s88fdxtnt3six\nnine33six15five6\neight3five26jdrrmvnzgxm\nzjzpg4onesixseven\n99fourfourtwo\nxlxmvkdkvrfrddhvgmhrqllvcpsnine2\nfourfour1xfpdnfqvcklf\nninethree7one2czbtzzfkz\ntwo41hgvsfzmtcz13\nsevenfour2sixq\nseventdnjmqhbmm3zgffshmjlzcjcvbl7mzc\nkhthzgjkzdkgjzxlkrvd4fivezq5\nfiveseven4onefourxtrfournine\n9onexh\n83ninet3drmkrtpdc6pmdrvdbx4\ntwofvxngqdhlpstlhfhxonegrrzfbsv7\nseven3fivethreepqzfl7\n56two229\ndjrthreethree3seven4zxlfntgssq\nthree3mdkbtxkkfour2gtpgcktqflrfs\n711q3ninevbsg2\nninethreenfone4eightseven\nonexfxnine64fivefour\nfoursqctmzlh64eight\n8one4mslnine65fnr\n9mbgn42bfhbfive79\n75kdhhfllf1onesevenssvztltg24\njxrpsflt52jz246five\ndlvhq4pzxcstjtxq9\nlgnxqnhjsbsixpm24two\n3sevenfiveltdmfzlq842four\nthreetwosix39foursixthreethree\nnine12bkhqklpfour3\n73534two1\nnhkbzfmgdhjjhggfivebncmnbskphsixnine2five\nxxzvlxmdtwo8threerbtwo1eight\n5bvzltnmqgnfqvqk3fivefhhhhnb\nqthree1fivettjdpgnmkcfive\ntknsqnjjczfbgshtvpdchpb6fngjztk77four\n1five2seven1rsqfsrxfcmnxtqj7nxzcc\nlcvzgqpmlgonefourseventwosixnine8\nssqcjzpjb5five5two\nseven7nrgglcjfour3two4\nfive4fivexnbnoneninengqngmg\n8fiveltmhtwo\njzmqtqbzh1\n1fourthree\n2eightfoureight2\n7nzkfvzfivemchbf\n7four2tpfvmjgdseven7\n1dzfpxmpxjscclshsixknfmgrpvptdf\nrm4618bkgpgmrhdldbftqzssbzsqxsix\nfivemmfxdgqpj7sevensh1nine\nqgqxznine19sixeightone\none7sevencnpdqmv\nvgm8four5ccrcnbddv2\n28gxkg245hpmd\nngvvdk41eighttwoseven\n1nine54\nonehcvdnt7twonflthree\n5nineseven\n8x\nfiveznsz56\nmplrzchnhjgxfmqbpfdm7kzrfour4threehjt2\ngdvjfqsqmnine67jlh9\n4tphmcfm33sixvgonefour\n61fourthree7ninethree\nfive4qsk42sixmxdrqcndcmvxcfqv\nrkrstzqone3bjjpgfrgg4859\n2oneqlsmvzznsxninefivesixksshvprpfive\n6fouroneoneqffour6jpfqpvcmx\nxshcfrdxnlv125nineone45\njqdvfiveglxjpbvzch9\n6jqmjmqznqmxkztdgzb\ntzjfdrzvzfivejsv49eightwozmf\nrslmsmj4vnq71eightfive\nvbmbfkfour26sz9\n6two769hffxnbnineeightnine\n9five3hqjdvxlmvvckndzhsix\njtknzpnbqpfgtdbnxkbthsix78\nfivethreerht48llrprvphts9cd\neight19\nsixone66six1seven\nfsgtskzsk9eight6\n2gb\nonecone79g8\nvgfqd49mk26qmxhqb\nlclctndpsffrbdmztb8\n1seven4\nthxj66ffbzzkz6twolzsbqseven\n2lrgrrhceight\n332three9\ndmtwonegppvjcmj2sqplmshfsd69\nfqdmgxtwotmzrb549six7seven\n87kkcbqvfftwomxsixsevenfrhlthree\n82l15vjddqclrtttsm\n2gjb3qhxvdrblq1\nthree4xvqb9nineone1fxsvj\n921\n4xzzdgzpklltssh95two\ntwoonejsmrddbtcng8zbhfbqvblbkmcpxst\n8five949\npmpv69fourddfxzrvphdsxz6\nseven6cnkkgphkdmvnxnxfvkfcmgthree\n98three\nthreert9rfsix244\n18jnzxfivefour1sixthreefjngk\n543vlxckcffkj\ncrxdmtzone1\nt99vqbpprvvzq3fz\n1mfourdvpfour5dzzdnjbdfxchkgpbnlss\nfive11six5\n9bbvmmxkfzldhjfghnlsevenfour3slnp\ntwokglseven1qqhsp57clctgc\nkmtrnineseven6one2\nfour1rg8five\nthreefour5sixsix\nmdvrfdcxndd1\nhqtsplgzqbrxbbstwolzrjcnqf3threeq\ndrhoneight7bnjh3fiveninejlqlljd\nfivefqqfcmzdfour1\n94thphvvkmkhpdnqsrh\nbsqhfjk3zxzhtqfncc7lcbzpgx\nthreezzjnlpgxhsfxmjl9588\nfourfourn279six\nfivevxzjpss5eight\n5fourc8fivenine9hmdgxjeight\n3threeninexjbzmsgrhkzxscfhddhsixtwohcgjfdszpxjsp\nsixknhkfsmqvvvone9jhfxjpr\n8bxgmbbdmrvpjfxbzd\n1dpxrdfvp3two8one8one\nthreefbvbfkk3scsbr\ngvmrjgjkdppdc9one\n133gglhlfprmtnqb5four\nvtzddchxktgsshpztmpxxvrhsjq1fkstwo\n8one5krgr55\ntwofive8nine5\npnjpkthree47eightjbsdz72\nsixonefhltgkbbbt8eightthreefour3\n72gbr5\n38mqgkbbkhbksqxfive61\n8ckhzckrkq88threessqqcpbnonejtxkxbbfive\n5threetwocrfsxsnf643\n9tlgcxkr\nthreehbmttm59sevenhfourhpqrjqj\ndfbqxtzkkpfseven42sevenseventwo\n86eightfvxgfjzrf\nvxfl4\nthree4bshflzvvdnbeightfqdlqnvlrkdxrtnpjcqn\n5sevenpxdrhmjbgdgpnseven\n2two4\n9xxnhc8vqmrnffmg3\n97kqfonefhhc4\n8fiveseven9\nnine6eight\n52five3one2fkmnk\nkfdxhkninestttsmtkmkhxbhhxlbxsix3pvdqd1\nfourninedhsvphbmhccjcktqsnthree8\nfivefive35three5four33\n4seven7seventhree6\nvfkjpfxltm5xcq5nine825\n52onegczcpghmzxvgjglbpjrtseven\n5zzlmvl\nkmsnldpssvdfive519kvmlnthree7\nlpkshkvrrl124dpgplnine88\n644seven9qdpbpxdlxf\n7jqc\ncbhpmvs7\ntwo7twoonebdmzftjjxrfmsix\n5mxxmdzdqdsixfourone\n3g6\nfglrhqsseventhreefivedpq4\nfivefiveglddjbfbk3two51eight\n7eightfjqqhljpxgvlghsc9lfqctddsoneninetwo\nmllcz8vczcgrbdn41qfglhzsqk\ngxstgzqfrfourgkxvnkpgckhnxj7seven\n93sevenjzvxc73\n63nine135gqdhbnine\nltcgbbxftssixdgtxdpp1kzbhvqt\n2z3seven1\n4two9qnfrvcbonesixzfr5\ntwojhlt32\n8one2eight\ncdjdb4vcvg98\nthree9nc\nqmfjqnsbthree97\nfivenine5one\nglnine55twoseven\nthree7k2\n437five\nsvzrgbljxdstwo9six6four\n383zseven2eight3xvqvp\n6fivegnglqj8tfkfgxkm573seven\ntwosixltnpgkkr3\n7sevenmphlgrdz72j\n5nine24\nhq9b4fourbgtjjgmzfj6\n1sevenczhtpmvmstvvgrbrhptxtlneighttwo4\nmbhfpxjbj1xdmvchpztmcxtlqjtk\n37hrqnnknjpcmmfhnbbmsps\n9fcsix51foursckzjtffvv\ngddqlqfourninenine65\ntgqvzvcnxs5lffnzhtd\n3bqfb8psxv38five\nfourcskck3seventhreefour8\nsix1sevenjdvjkrn6six88\nthreetwothreeone87jlhgbrtgvdrt\nsevenmsldjlpvzhtwosix4threefour\ncthhqm2oneeightfour\nrpbbnszfxg8ninetbeightfiveqskfzcf\nsbjghdfgheight9seven\nnfj59pjhrzfourfive\ntwotwoseven6nxxblhkdxct8\n5lgxnqdfiveninedtdonespkvnonelzphcrxr\n7qkphfqbbjvninedprb8cgmpcmsz79\nfkmhvlvjgczxsllskzdone8\ndzjvft8fivecpclmdtwo5453oneightt\n6bjbvjgm\n64ckqzvnjdhnine\nfour5eightwon\nsix3vggtzjxtsjfpbqh\n7one5\nonezccmbfq3\n7gseven3seven\nnxsevenone6ninep\nhmmz2hvmxcmqtsrxdllrnzrgzcx85three\nmskcmjfdvxmxlv5dnhhkvhx\nkrxjnz48sixzpqzppcngmsixb\n7threedgsgseveneight\nnfhpdbs4988\nfournnfsjtwosixtwo2fourcnvdgmcfive\nptfhmlmjzltftmn2nvxvksbqhl7two7gdxb\nthreehqgqzx32\n598flbdmfivenine\n1tvrxqvsixfour\nthreedz8fourthreencxs\n8foursxvfourlrgqjh\n31vjeight2nfgkjvjqseven\none2lvvlseven9\nfournine5xtpl\n2fgfbjg\n2onefour\none3vbrbkjmkqxs9t\njdgcvgcfsdbhvmeightkf8onenine9\n2onekqfmjcftms\nseven68two4zn\nsevenvqpmhtxhncvff5one\nmfqgljnmonepldmbnqphhtdonefive3xr\nrrxkcnineninebsh3sixnine\n81ztjznine\nmgkgddsonethreebpmhnpfnqjjhzthreehdpnzmht55\n3five3\njeightwotwobqsbczrlfgj7\n6pvghfqtfjc\nthree3onenine4pmtfive\nrnrdbjxtwo7fivesevennlpfrcptqtwofiveone\n7psthptpvbh2\n4bvsgkshslg44vkfour\nthree1tworkxhhps\n299one\nrghnkhlxqz91\n76twocfcrone3\n97dkkfonen6jnsmlh\n2hnqhcpr4\nseven7llhskvhnxz4\n6ffjvvdgbtdxpmvdsk5qvppxtgblpvl\nbtsl8threenvhjkcckkgfsfnjgthree49\n31vvmjfz\nhjbkfgnmk35four4kgxb3two\n5zkljxdhzpcfb7five8zpsckrfkvbeight\n8onefourtvdgninenfive\n6eightzcvfourrzqp7\n3foursevenxjkmnnvqhq\n5274eightwohx\nvktjgvnztldsfbnqgqt9seven35\nj5xqqltd3fouronesix\n4bnpgeightsxlrbxdsnbdc\n3zzxtcs8\nddjd13zpfqqx\nfivenmninerrzkdp1\noneone989bn\nsixfltmpjjqnmfpln5\n5ninedn\n562\n47eight8hqsvhpk\nnreightwofjnkrsgcdnt75\nfour1qvdfsfthgbb3kjrlh\nshtf21\ncqdlfptg2zzninelsmgpvqvrqnz7\nshsjvztwonine8fourgrtgbdfdlqjvmktnrt\nxpjzeight9rr\n56fourcjrnmttqp8two3\n3sevennxshxncmcfnxpv\nfourvjgpcsqfjmtwofour9nine\ntwogjvtlpnsmcjqkgkkj7fivefmmthrfs\nfour9kssrrzrlbqeight\nthree4eight1xzrqjpxpnxqptnqvmqsdj\n6phjpjd\n5twoklcnsq1xcone\nseven1bvnpqqkqcmzdllhj23five\nff19sknzln8jrdlhptfourz\nrdrtwonehvdxccpftplzfqpzcrjsix6pfr\nseven3llthreeshcb\nqqzlgseveneight6one15four\n7pzmfpvzdhhnxpsxd4vrqbmgtxjcsvfxdnxjzsevenkd\nthreelvkggjfivemlddrfrrllgfssfmgs17\ntcbckd2\n7four7krmsixffv9threethree\nglvknsmmgn2onexjdskcqq3threeeighttwo\n1fpt3\nqddscqpjlpnnlsneightkm6sixnine\nhppd1xlkgsmtfxxbvcgzbdfournine\npqdmrhvxdqbrkgjdc36mbq\nrpllrjvrvfourfivefour3\ntwo5neightsevendmp1ninehjfvvplmfc\ngppnq7\nfoursix81onefkhzshsfxlxgxjzdd\n13five\n3sevendjlgttvnsevensjpgsjt7\n172\nthree69threefour\n3tcnfpzpkz\n5vxzzfhg24bznvrmsjmtfmfjgxdseven\nszhmlqxfffiveeight71nine\n8five89zjvjqtpmlqveight5two\n5twopjlvzlrltv\nsbjtbfqkshrrtlrxvvq97jrflnxcrnfive3\nqhoneightfoureighttwo5xssdqmvljfivenine5\n3eightsix3five9\ntreightwo3sqzninenine1sxk85\noneeightxklmdbfkqjfour9threetmlnine\nlxfxplrqeightninethreeppfnzfvndjn8qrmsjq\n5one55eightwog\ntworlqkztqblj9ltxlshrhjlttnqzhsdbpzzzbxdn\n53vvvsmxjvs\n4nine2qpgvdtxldq4four8jnf\nkeightsevenhcvrpztgcvsxszqdlzd1threenpvszb\n89eightvdvklpxfntwo\nndtqsrfjmtzgdptlp8pnnf4onelq\n6threesevenfx\n6fjbglhltcfdxcbnzlcrtwoneqx\n3jxtvj1onemxbqxlnm6sevenjzkhljktcm\ntwogkljgjbxzhlnrdxkhqvsevenqqttpjnx3\njmdcbcllcv1\n3twosevenxgtbzsfhrbxsgk9sevennine\n372\n52qzldrtxqxnfivekkdqqs54\ngcmnp51hxhvdbk22gdvdfour\n6bpdjtwoqmqtvzqfive\n4zmfthreefivefoursix1csnrznptgr\nhrdpbqfonefjqbbjrnine9onenine1\n47bdtqmbdvhmsf3five86vhclmd\n295two7four\n3onenineseventkgxthreetblbmljhkhfour5\n4four8fourthreek36\n5rqnsvdhkbrvrltcqpsd\nsix4fourfourthreefour\n4967\n99\n71seven9kcmclngvpmlkkfngssveight\n79eight9six\n3sevenp\nseven69lxsllgjqdvfoureightfive\n8six5grssfgfr78two\ntwomkpjc18fkkhlnkgkxfjnine\ng8qcq9xqlzgzmmrf5vhm\nzsix8ssfrjqbnztttvszlsqeighteight4slzcrnkvf\ndmfive1ffprhgfnrx49\n7cfdcptvvnvhfivetwo\n18kqkxdlnhjbqoneszclkzljhnhqd\n52lcclkxsix\nthree7five5cjtgpzjfoureight\n53pchxqfpcfddsqzl7oneeight\nctjg84hqsjxfive\n1znmjftpq7xbxt\nnine71krj7fivenine\npgr6lfbhkpsjhhjpgqnxtxmkpdktnine21\nxhkeightwo14four23eightqxhpct\nx1\n7357fivehfbbgdhzg\nrktbdvm41seven\n6fshmcmxdsfgxrcfvone\nnndrdq7six\n5two7r\none78gnsdkhkqfourfdsghnll6g\n3sixht\n96mbxlpgctgf1sixfivefour\n2sixbcqskzpqbqtt8sixrtngznb2\ntwoone9fvmczldvtk\n4jpfmgzgdhs\n3ninefour\n61nine\n8kleightoneightrhf\n51sixg\npckfpxzmrthreejvh8eighteight5\nssmgkkblkfive68eight2\nnxkxdtckfb95bpdtrdhhltjdqpnine6\nkqjkvndr3lfqthvmbrbfgzrbcntpcrd\n44sixtwonineonefour\ngtvlneighteight5961nine2\n14lzxzbvbg4sixeightpthgz3\n8814eighttph\n7rpf8fivedkfivesfxgkjmniners\nsix6zsxzcfmcxzn2\nfour12dkgnmsjqtjp7zmzbsp\n8two69six\nthree13six4three6one3\nthreedfmcmxgxzhglv36five\n8kkjrbkmbt2\njhttwone3fivethree8ninehjgnvmxtkdcpmhhvrb\nsixtwofourseven86zs8\n26seven\n28three6\nqpsznsqqfb4ntcdf\n9zbqnvnlgrjcnpmbkdmtvmdfouroneqctnnm2\nonefivekpfrdbxcmn3sevenonerfzmcsvbjgtcxndv\nseven814nskmxzpcnvfppntrssthree\n74jfmrfkznsfiveeighttwofivethree1\nsixgh8\nfhnlhfour3jhrzjdsrfour4five\nqgzblcqp9\njrk377gdrjfldpqmeightffrlxffive\n36eightsixgvninezkl\n3hshcq9r\ndmrhgkv2njtwo95\n6pktmsdxpmpvmpfc4eight\nfive1five6hnpbzpmcccnqpfive\nlmsbmvmrrljf4ninesix4threertxtsvfour\n419six\ntstsix3six\n37seven\nsixtwovghcj32pj\nx66four6\n72twosixnvdcq2nvkmgbsb\nkvxprh1tnlrpmcldsshh5kcdfxmbcbfour\n2threekqqcvtxjrktkqpddqhfourpbqqcrgbdeighteight\nrltdhqvnm8\nsevenninetjvvfxcv7\nbhmkxpvhrqkfive8fourtwotwo1lcbkk\nthreepcnnzhkbzfbhqtn3eight544rtsrtcc\neight7onekqdvzvsmxcpgfjlllhtptnrmv\n47zsqd2pnsgjhpkxljpvnine\n8ldxnsfour\ntwo7vpjfvhtmqdcrqlonekkqrnc32\n12znrsmjqf\neightsixeighttjjhqsevenns1qbkstj\npcpcrqshftwo4\nsixtwoninenmffcmfldb81\nfive75pvngkvx9nnlttwotwonev\n1739lkzpsfzdsixfoureight\nthreekqvbcfiveltbrtjjtdcchtjthxvztkjpd3eightsix\nqrxtwonefive9\n96zthtvk8228bmfllnc\ntmnsgdz3fourtpmtcvtfqdcteighttm\nrmpbtfqxmmtwokn5two\nninefive65\nqzmthdvvzmxthf82fivesevenbvdt8dscnl\nonenfhgqqpdrtxvzn961sxjhxlg2nmxzkbp\n9eightone\nfourfhnzmcfour6lshj\neightzlffkpcbrbjjpmjrz7tcgjonefourthree\ntwomdzv77four2437\nbvcgtrxthg863\n4fpmpsxzdshdszvckpxqkgfourxnrpp\none3sevenmfmlk\neightthreenvfp439\nbm9\njssc74tlhvfdjs6kdxzrskxrdtfqnnkfczr\nv2eightsix\neight2mflbkmeightqfivethreefour\nfdchmls7\n7cgtkzkdvkfour\none237\n69tfpdbmcrbqq1\nponeightfiveninefivetwosccxmtbf28\n87sevenpzbjhjvqonetwo\nxsbsqpdlp9threejpjmrctbkrjrsevenone\ncmhxdcpl85jhtpnttwonzlcrcfivevxxqn\n1zxccxdpkl9tdqkgfxsevenvhlrmvfhvvzdzbk\nlqmnvtdqngqjh355\nmcsttwo43bsdrtt7three9bhlh\npvmtwonefour8eight8\n2tgcpztjnqkldhshbmxnm\n5zqkkcnnsh3\nvqfmjzln6xnrpnrpqkdpqfgxfnine2\n87eighttwo1\nsevenxrdhnpz87zzjzhxcbbxqxeightsixoneightssl\npdpgbbhpzzmhqfhkprgx1twoonethree\ngggfnvtxvpsix6slxrninenine56\ndxbhdfrvgqkhch5seveneight\n7five3\nfourfive1fourptggnine\nthree6npf45kjrbrg94nine\n4twonev\n4rgpf98two\n7chcnzghmcznftwo5\ndvskzxcgnljmgxfntwo73\nqboneight5nine1vftwotwo\n54lqlzgjqctwosix\nzxqnine57\n3onethree\neightfive2xxxtwoblfjhlq9b\n6ninesqvjsrz813\n6734\n2jcxkckrj7crncbonevskt\n167fivef1\nthree9five2eightfivetpdlpxljgnineone\n9two6five\nxxktwoone2crphztrtndrjgs\ngcjzskxzmtwo45seventwoqvzpfive\n8phdzrrlkhqjrmlgshlmonefggsnslj\nthreer3oneeightfive43\nsix1one\ndzoneight8\nsdztwoneghxmqxlgv68\n9onecjone9ppvlt\n2svml3fourfl6five\nclknqq83ninefive\n1tgfkdnqgoneightzf\n9onesixfiveqcclnl\nonedvchs4\nstrmhlxs2threefouroneseven\ndpbxhlgkzbtwo32nrdtwo\nseven38sixfivehlkdsj12\neightfive6\n8sevenfour8mknvzhfnzvbjplldpqkvvrcrmfourfive\nxmzjtktmsixsix2ninesvqgxc\n7fourbpn1four\n8nineqksmfivevdknb\n8kfv4ldzf\n71fiveeight\nbbmtpsdbnktndcnthreethree2jeightsevenxl\nntggphnine3\n6vfb5\n9pqjp3fivedssspd\npmsixeight7x6\n1nine5sixoneightmm\nthreefbhsjcqrstdzkfgsllpktqljnlsevengbvpmvhlthree5two\n9rqfdgxl4pgfbdsfonedbmjqthnsevenone\nccqknlg29bzxdrdx\nvtbjqtjhr23\nsevenvstvhd1ninesix\nsixdjszthree7\n6jhnj4\n2fivenjrfkz\none55\n69fourvttjxr\nkqpqtkss7lxfxsfsmxqkfpcrjm\n459xr\n46nzsfive\ns4jhbqgpmsninesixfivefiveone\n4twofivefive\n528nmtcfptktmfivesixfour\n78qnkp\nzbgrb4mskvlrkcftbpjr2twones\ntvcjhbv4nine\nthree6mxsxvlljdnsgtzrfouroneightlp\n6jc7two\none9kmkrvjrzbv6eight1seven9seven\nfivesixstjrxhn7zmjthsdk3fcjsq6\nncbzqeightfivetbpmzrvp44ctpjhg\nsevenrgvgkzktvjdd92\n5eightsj\n83nine7hzjvjxmbn\nsgspsggkbzfour9\ntwo5fivetwofive41eightwosm\nonejmk5eight\nnine11\nxlhx53qrnvqfiveclj7six\noneeightone1five8one14\n56d3\n5sevenseven9dnhkhhpjzbqkrt31\nr6one\nsevendh46\neighttwofivesevennine915eightwomjk\nfour26three\nsixrspdrhltrseven3\n3bf1two\n6ts\nsevenhnvgvkhztsqbqcvm17btpzrrrfpsnr\neight8fourldzrzhpmppkcfcchtzpteightthree5\n124\n379qtfive\n4pp\n128pqgtqrjdhlrfivesixthree\nf7ninesix\n16hqkxjnmd\n9sixsix98qhblkzh\n13nine8glvsfkrqs\nninegdnhsfgckc7two\nonenine69lzqx7gxhcdcg34\n67993fivexcx\nnine6dbqrthree9five6bone\nzvfmz6thgmqgqsvnqpxtdx575\neighteight2\nfive2five\nlprqsrxqceighttwo744mptdngcgmchkg\n1seven3ztfhcsrkdj76rmjz8\nh4fiveeightfournineblvhhnnqpsix\nsixk4\nnine7vhpkkztwo\n8jfrfchzpk54sixcfbdfrccpv3\n9zkxnqjfzcrdseven8ninefivenine\nninefivenrfmttdq9vgtnqsxcmh\n5tgjvsix5\n722fvzzn\n23threethree\n7two5hszrj5nineonethreedp\n2lpdzpzeightfive2onethree\nlptvrd5qjnstntpld44eight6bflqd\nldh1three\ntx9kgbrrvthreezkxhpn6\ng5gks6fivel\nhlblkd3threedcqhbtmktvtzqrmjhgtgtxnfbbdffmckbqxxp\ndfjmnine71vfndzz\n66eightfivepzqvsrlrr5\nseven2threeninemcmlhfmkndhtdeightthreethree\n6ktfgglmnvjgxsevenfourjzgp6eight\n76zvbblfrthreeeight2ninefive\nfivetgprhplmsmhslrtpr31six4\nsevensevenone71rgqb6three8\n34htvldclbxmeightwon\nfive1klzmf3\n5dqpfivethree4\np6\nxrtwone719\ntmtwonefive3sevenqqpz8twotdsn6\nqvcdkggrh644\n899tnqhhzcjxxllmp139prgx\nsixnrsgvmrggcl6\nccjq3one\n5threejbhlrrlone1four\nninefive1ninevhndvmrbqzcnlccptbvl\n3sevenonepnjxhkbtc\n3ninelck34eight\njmjcc88trmxvsvkgxdjlh78jjtwo\nmpfpnqqrsspmdjpkmrrlljrlsddnsix38three\nmtttwo97\n4xfnjbsplqnxtns13qlsnbhn\neight69eight\ntwo6dtrfmgh4\nqtgvqlzp175zxt5\nsgmqvl9tmnthree2\ndb47six6nkgdqvc\n6jrsixlldvpn2onekrqninejh\ngdrqhbkpg3one\n7threeseven8threegqlml6mjmtc2\nnptwothree8five\n8kxmjpnctddeightnine5twonevzf\nthreefour7fivelhvxvhlqzfn\nspfgjxkcnfourllmv3v9four\n7xzksddd7v8pjrv4five\n68fq1four\nthreel6dd7one\npnrsevenseven5four95\n4sevenqlzjjcdsvfdtzcjp\n1kxfmrsix58ninenrmmlvtd\n743threeqfxmmtvkczhteight\nbslvhvlspr8dsmhjlrcfjsixnine\n2mpjmtrrnineeightklxflnvgr\ngxnzcrf4xvgcgb9one368\n4three5j28xl\n2seventkdnfrkvnd7\n7llbkvzblfdtwo667six\ntwo7onethree2fivesixhgjg\ngthreelprkltjbqp7sevensix\nq41cpcpjh\nnine42fivefzztjknfv\nfive3kblrgsonedxjhsnxfour56\nfckqjbs1cbblpbvtxdlcvtwotwo\n9xvjszgkbcfourninerp8nine\nvmlx27oneqzmffghcmfpn47\ntjschpthreecbb6sclbfivejzlbd\n77sfzkxdfdjnbnseven3knnqgvjbm\nthreeqgfcrn2\nfour7qcrxlcjsklhjcsz2731\nsixreightsixqhflvseven8\n1xxngrjnfsixfiveslrfivenkplbgnqxg\nonefnhbcjqrthree7jlbcfgdl5\n81oneightrgr\nsixmqgdfgscr6three\nsixxlcsix2\n1zrqllfszlhqgp2\npcvrlxfml5znrjc542\npctzcvgrfj82\njnrlqsptfivetwokghkr8sevenlzdf\nkfbdrtgtgsix9sevenone\nppzqjzvvfive39eighteightthree\nbzblbmns62seven7jvdgxknx1one\nfoureightsevenmtwo7\n3dxbt6h\n7phdxhtsbmfzmhzlsh3four\nfsrcjfour2kllnxdbfhnine3six\ndp6sevenseven\nnqf3pv7\nfivesckjvndbf1eightfivejqzcsn\neightqxbx3jdzmcsvrpv9fplreightzdrlbpfj\none6twoeightkrrfdgrttxvthgfour\nsevenseven3sixsdrtj\nrstsevensix4oneseven7mrccrxmht\nfourthreehjthr8five\nfourxmxncslkq71\nseven1b5two\n3four9eightfour34\n76one6four\nsixsixpzghqfqndvpcmzt4tftxxpjbghs\nfouronenine7eightwot\ntwo77mplckblrclfqgpgsx24seveng\neight6twoeight18fczfn5four\n7dxqshsmhrvbnfgjq5\nnine52three\n3thbtwo\ntwo9sixeightoneeight\nmklp1qqgkgcceightfour77\n13onermtxflmnmq8qxhkhxthree2\n397rhmknine96rzfgbr\n8fourseveneighthb67\n89625t\nfbfxgqsqhthree74seven5\n8threegnqhhpfx\nzdsfs682\n76lqkvfhtdpseven4lfpvkxjgqtwobbrz\n7zthreesevenfour9four\nninefc6eight4two5oneighttfn\neight32\neighteight6lqgrgbntgkkzfqdjhxtwo\n8vdzgsqnsix551\n6klxfrfqzbsrnxmmdbnqbvfpzcjtx5two36\nnhcdmxpvg5kmmknrnine\n17tpsvtclthree\nfive84eightfjrznfsrb\nfjzfb6onefourhtlmvlns\nzqmmfdkl8\n3sixeightnkgpssqnkrsclmshzzfhxxhvxlsljgfgnltbpc1twoneqd\n8frbccqkvdtwoc\nzeightwoeightptmkdhx3eightnineqx4\npbfnine6three8six5jtxmeight\neight78\n1bhlgn5five\nxnhhlgfrqbgfhhnvllhptfh3\nljjllzbbffpxcjrztzthreermg6fzqqpd\n72mmjrfjvlzone3threethreesix\nfiveonecfsfsix74twocllbfnptkgttf"
	input := strings.Split(rawInput, "\n")
	want := 54076
	result := calculate(input)
	if result != want {
		t.Fatalf(`received %d, expected %d`, result, want)
	}
}
